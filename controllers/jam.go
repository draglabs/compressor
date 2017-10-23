package controllers

import (
	"compressor/db"
	"compressor/models"
	"fmt"
	"io"
	"net/http"
	"os"

	"gopkg.in/mgo.v2/bson"
)

//FetchJam func, fetching a jam by
// a given id
func FetchJam(params *models.ArchiveParam) (*models.Jam, error) {
	var jam models.Jam
	ds := db.NewDataStore()
	defer ds.Close()
	jc := ds.JamCollection()
	err := jc.Find(bson.M{"_id": bson.ObjectIdHex(params.JamID)}).One(&jam)

	if err == nil {
		extractRecordings(jam)
		return &jam, nil
	}

	return nil, err
}

func extractRecordings(jam models.Jam) {
	for _, v := range jam.Recordings {
		extractURL(v)
	}
}
func extractURL(rd models.Recordings) {
	DownloadFile("temp", rd.S3url, rd.FileName)
}

// DownloadFile func, fetches the s3 url file and saves it to disk
func DownloadFile(filepath, url, name string) (err error) {

	// Create the file if it doesnt exist
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		err := os.Mkdir(filepath, 0700)
		return err
	}
	out, err := os.Create(filepath + name + ".caf")
	fmt.Print(err)
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
