package controllers

import (
	"compressor/archiver"
	"compressor/db"
	"compressor/mailer"
	"compressor/models"
	"compressor/uploader"
	"fmt"
	"io"
	"net/http"
	"os"

	"gopkg.in/mgo.v2/bson"
)

var currentJam models.Jam

//FetchJam func, fetching a jam by
// a given id
func FetchJam(params *models.ArchiveParam) (*models.Jam, error) {

	var jam models.Jam
	ds := db.NewDataStore()
	defer ds.Close()
	jc := ds.JamCollection()
	err := jc.Find(bson.M{"_id": bson.ObjectIdHex(params.JamID)}).One(&jam)

	if err == nil {
		currentJam = jam
		extractRecordings(jam)

		return &jam, nil
	}

	return nil, err
}

func extractRecordings(jam models.Jam) {
	extractURLAndDownload(jam.Recordings)

}

func extractURLAndDownload(rd []models.Recordings) {
	c := make(chan error)

	for _, r := range rd {
		go func(d models.Recordings) {
			err := downloadFile(currentJam.Name, d.S3url, d.FileName)
			c <- err
		}(r)
	}
	for i := 0; i < len(rd); i++ {
		err := <-c
		if err == nil {
			fmt.Println("no error downloading files")
		}
	}

	archiveIfNeeded()
}

// DownloadFile func, fetches the s3 url file and saves it to disk
func downloadFile(filepath, url, name string) error {
	fmt.Println(url)
	// Create the file if it doesnt exist
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		os.Mkdir(filepath, 0700)
		//return err
	}
	out, err := os.Create(filepath + "/" + name + ".caf")

	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	if err != nil {
		return err
	}

	return nil
}
func archiveIfNeeded() error {
	_, err := GenerateXML(currentJam)
	if err != nil {
		return err
	}
	if err := archiver.ZipArchive(currentJam.Name, "archive.zip"); err == nil {

		url, err := uploader.Upload("archive.zip", currentJam.Creator.ID)
		if err == nil {
			mailer.SendMail(currentJam, url)
			uploader.CleanupAfterUpload(currentJam.Name, "archive.zip")
		}
		return err
	}
	return nil
}
