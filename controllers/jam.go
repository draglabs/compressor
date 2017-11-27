package controllers

import (
	"compressor/archiver"
	"compressor/db"
	"compressor/mailer"
	"compressor/models"
	"compressor/uploader"
	"errors"
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
	err := jc.Find(bson.M{"_id": params.JamID}).One(&jam)

	if err == nil {
		currentJam = jam
		err = fetchRecordings(jam)

		return &jam, err
	}

	return nil, err
}

func fetchRecordings(jam models.Jam) error {
	var recordings []models.Recordings
	ds := db.NewDataStore()
	defer ds.Close()
	err := ds.RecordingsCollection().Find(bson.M{"jam_id": jam.ID}).All(&recordings)
	if err == nil && len(recordings) > 0 {
		currentJam.Recordings, _ = setUser(recordings)
		extractURLAndDownload(recordings)
	}
	return errors.New("Not enough recordigns to process the zipping")
}
func setUser(rd []models.Recordings) ([]models.Recordings, error) {
	var usr models.User
	recordings := rd
	ds := db.NewDataStore()
	defer ds.Close()
	err := ds.UserCollection().FindId(rd[0].UserID).One(&usr)
	if err != nil {

		return recordings, err
	}
	for _, r := range recordings {
		r.User = usr
	}
	currentJam.Creator = usr
	fmt.Println("fb email from setuser: ", usr.FBEmail)
	fmt.Println("creator email", currentJam.Creator.FBEmail)
	return recordings, nil
}
func extractURLAndDownload(rd []models.Recordings) {
	c := make(chan error)

	for _, r := range rd {

		go func(d models.Recordings) {
			err := downloadFile(currentJam.ID, d.S3url, d.ID)
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
		fmt.Println(err)
		return err
	}

	if err := archiver.ZipArchive(currentJam.ID, "archive.zip"); err == nil {

		url, err := uploader.Upload("archive.zip", currentJam.ID)
		if err == nil {
			mailer.SendMail(currentJam, url)
			uploader.CleanupAfterUpload(currentJam.ID, "archive.zip")
		}

		return err
	}

	return nil
}
