package controllers

import (
	"compressor/db"
	"compressor/models"
	"io"
	"net/http"
	"os"
)

//FetchJam func, fetching a jam by
// a given id
func FetchJam(id string) (*models.Jam, error) {
	var jam models.Jam
	ds := db.NewDataStore()
	defer ds.Close()
	jc := ds.JamCollection()
	err := jc.FindId(id).One(&jam)
	if err != nil {
		return &jam, nil
	}
	return nil, err
}

// DownloadFile func, fetches the s3 url file and saves it to disk
func DownloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
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
