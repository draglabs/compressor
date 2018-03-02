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

// Fetches a jam first by JamID, then by UUID
//
// Params: a JamID and an UUID as passed in from the API call, held in models.ArchiveParam
//
// Returns: a jam model and and error; jam model is null if no jam was found with the ID given
func FetchJam(params *models.ArchiveParam) (*models.Jam, error) {

	var jam models.Jam
	dataStore := db.NewDataStore()                  // contacts the database and retrieves an object representing it
	defer dataStore.Close()                         // closes the dataStore instance once the function returns
	jamCollection := dataStore.JamCollection()      // fetches the jams collection from the dataStore
	err := jamCollection.Find(bson.M{"_id": params.JamID}).One(&jam) // searches through the collection to find a jamID and assigns it as per the models.Jam to jam

	if err == nil {
		currentJam = jam
		var user models.User                        // sets user to represent the model
		err = dataStore.UserCollection().Find(bson.M{"_id": jam.UserID}).One(&user) // looks through the same dataStore to find the user by his ID and sets user to that data
		currentJam.Creator = user                   // the current jam's creator is the current user
		err = fetchRecordings(jam)                  // fetches the recordings, the function down below

		return &jam, err
	}

	return nil, err
}

// Fetches the recordings
//
// Param: a model of a jam with data inside it
//
// Returns: an error if something went wrong
func fetchRecordings(jam models.Jam) error {
	var recordings []models.Recording                                                       // an array of the Recording model, initialized
	dataStore := db.NewDataStore()                                                          // calls the database for a new dataStore
	defer dataStore.Close()                                                                 // closes the dataStore once the function returns
	err := dataStore.RecordingsCollection().Find(bson.M{"jam_id": jam.ID}).All(&recordings) // finds the recordings for the current JamID
	if err == nil && len(recordings) > 0 {
		setUser(recordings) // if there is no error and if there are recordings for the user, run function setUser, with the recordings themselves passed in

		return err
	}

	return errors.New("Not enough recordings to process the zipping")
}

//
func setUser(rd []models.Recording) {
	var recordings []models.Recording
	dataStore := db.NewDataStore()
	defer dataStore.Close()

	for _, recording := range rd {
		var user models.User
		err := dataStore.UserCollection().FindId(recording.UserID).One(&user) // looks up the UUID and modifies the user object
		recording.User = user
		fmt.Println(err)
		recordings = append(recordings, recording)
	}
	currentJam.Recordings = recordings
	extractURLAndDownload(recordings)
}

//
func extractURLAndDownload(rd []models.Recording) {
	c := make(chan error)

	for _, r := range rd {

		go func(d models.Recording) {
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
//
// Params: filepath = the filepath, currently unused; url = the S3 location; name (string) = the start time + the user's first name
//
// Returns: error if something goes wrong (read/write error, file not found)
func downloadFile(filepath, url, name string) error {

	// Create the file if it doesn't exist
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

//
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
			uploader.CleanupAfterUpload("temp", "archive.zip")
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
