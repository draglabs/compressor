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
		var user models.User
		err = ds.UserCollection().Find(bson.M{"_id": jam.UserID}).One(&user)
		currentJam.Creator = user
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
		setUser(recordings)

		return err
	}

	return errors.New("Not enough recordigns to process the zipping")
}

func setUser(rd []models.Recordings) {
	var recordings []models.Recordings
	ds := db.NewDataStore()
	defer ds.Close()

	for _, r := range rd {
		var usr models.User
		err := ds.UserCollection().FindId(r.UserID).One(&usr)
		r.User = usr
		fmt.Println(err)
		recordings = append(recordings, r)
	}
	currentJam.Recordings = recordings
	extractURLAndDownload(recordings)

}
func extractURLAndDownload(rd []models.Recordings) {
	c := make(chan error)

	for _, r := range rd {

		go func(d models.Recordings) {
			err := downloadFile(currentJam.Name, d.S3url, d.StartTime+d.User.FirstName)
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
	tempPath := "temp"
	if _, err := os.Stat(tempPath); os.IsNotExist(err) {
		os.Mkdir(tempPath, 0700)
	}
	out, err := os.Create(tempPath + "/" + name + ".caf")
	fmt.Println(out.Name())
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
		fmt.Println("error from generating", err)
		return err
	}
	if err := archiver.ZipArchive("temp", "archive.zip"); err == nil {

		url, err := uploader.Upload("archive.zip", currentJam.Name)
		if err == nil {
			mailer.SendMail(currentJam, url)
			//uploader.CleanupAfterUpload("temp", "archive.zip")
			addURL(url)
		}

		return err
	}

	return nil
}

// Add URL Link back to the jam
func addURL(url string) error {
	store := db.NewDataStore()
	defer store.Close()
	err := store.JamCollection().UpdateId(currentJam.ID, bson.M{"$set": bson.M{"link": url}})
	return err
}
