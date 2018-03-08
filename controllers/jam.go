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
	"time"
)

var currentJam models.Jam

// Fetches a jam first by JamID, then by UUID
//
// Params: a JamID and an UUID as passed in from the API call, held in models.ArchiveParam
//
// Returns: a jam model and and error; jam model is null if no jam was found with the ID given
func FetchJam(params *models.ArchiveParam) (*models.Jam, error) {
	fmt.Println("Running FetchJam; JamID: " + params.JamID + ", UserID: " + params.UserID)
	var jam models.Jam
	dataStore := db.NewDataStore()                                                          // contacts the database and retrieves an object representing it
	defer dataStore.Close()                                                                 // closes the dataStore instance once the function returns
	jamCollection := dataStore.JamCollection()                                              // fetches the jams collection from the dataStore
	err := jamCollection.Find(bson.M{"_id": params.JamID}).One(&jam)                        // searches through the collection to find a jamID and assigns it as per the models.Jam to jam

	if err == nil {
		currentJam = jam
		var user models.User                                                                // sets user to represent the model
		err = dataStore.UserCollection().Find(bson.M{"_id": jam.UserID}).One(&user)         // looks through the same dataStore to find the user by his ID and sets user to that data
		currentJam.Creator = user                                                           // the current jam's creator is the current user
		err = fetchRecordings(jam)                                                          // fetches the recordings, the function down below

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

// Sets the user of the found jam to the one in the list of recordings
//
// Param: an array of Recording models
//
// Returns: nothing
func setUser(rd []models.Recording) {
	var recordings []models.Recording                                                       // initializes the recordings object, a list of Recording models
	dataStore := db.NewDataStore()                                                          // calls the database and retrieves and initializes a new dataStore
	defer dataStore.Close()                                                                 // closes the dataStore once the function returns

	for _, recording := range rd {                                                          // loop to go through the list of recordings
		var user models.User                                                                // initializes a user object as per the user model
		err := dataStore.UserCollection().FindId(recording.UserID).One(&user)               // looks up the recording's maker and modifies the user object to be that of the recording's maker
		recording.User = user                                                               // accesses the recording object from the list and sets its user to the one just found
		fmt.Println(err)                                                                    // prints an error if one exists
		recordings = append(recordings, recording)                                          // appends the new recording to the current list of recordings
	}
	currentJam.Recordings = recordings                                                      // sets the current jam's recording list to the one just created
	extractURLAndDownload(recordings)                                                       // downloads all of the recordings into a local folder as found in the list of recordings
}

// Extracts the URL from each recording from each element in the list of recordings, downloads it, and gives it a filename
//
// Param: an array of Recording models
//
// Returns: nothing
func extractURLAndDownload(rd []models.Recording) {
	c := make(chan error)                                                                   // makes a communication channel for all threads so that the program doesn't crash before all have downloaded

	for _, recording := range rd {                                                          // creates a loop to go through each recording in the list of recordings
		go func(rec models.Recording) {                                                     // launches a goroutine
			formattedStartTime, _ := time.Parse(dateTimePattern, rec.StartTime)
			err := downloadFile(currentJam.Name, rec.S3url, formattedStartTime.String()+rec.User.FirstName) // initializes an error object for the downloader, gives each file a filename
			c <- err                                                                        // tells the channel that the download has been completed, and sends with it an error
		} (recording)                                                                       // does something, maybe runs it?
	}

	for i := 0; i < len(rd); i++ {                                                          // checks through the channel to find errors by iterating through all objects in it
		err := <-c                                                                          // defines an error
		if err == nil {                                                                     // if there was an error
			fmt.Println("no error downloading files")
		}

	}
	archiveIfNeeded()                                                                       // archives the file if it is needed
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
	out, err := os.Create(tempPath + "/" + name + ".caf")                               // adds file to "temp/name.caf" and appends ".caf" to it, disregarding if it may be a ".wav" on the server
	fmt.Println(out.Name())                                                                   // works by going through all recordings in a jam, getting the recoding name, which is "rec.StartTime+rec.User.FirstName"
	                                                                                          // the url to the file, and replaces the url's file name with that as above, and appends ".caf".
	if err != nil {
		return err                                                                            // exit function if there is an error
	}
	defer out.Close()                                                                         // close output stream once the function has returned

	// Get the data
	resp, err := http.Get(url)                                                                // downloads the file

	if err != nil {
		return err                                                                            // exit function if there is an error
	}
	defer resp.Body.Close()                                                                   // close output stream once the function has returned

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	if err != nil {
		return err                                                                            // exit function if there is an error
	}

	return nil
}

// If the file downloaded correctly from the current jam's list of recordings, this function
// generates the XML file for it and appends it to a zip file containing the archive.
// It then uploads the zip file and sends an email to the user with a link to that zip file to download.
//
// Param: nothing
//
// Returns: error if something went wrong
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
//
// Param: the url
//
// Returns: an error if something went wrong
func addURL(url string) error {
	store := db.NewDataStore()
	defer store.Close()
	err := store.JamCollection().UpdateId(currentJam.ID, bson.M{"$set": bson.M{"link": url}})
	return err
}
