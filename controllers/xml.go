package controllers

import (
	"compressor/models"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math"
	"time"
)

// Header is the header of the xml
const Header = `<?xml version="1.0" encoding="UTF-8" standalone="no" ?>` + "\n"

// HeaderDoc is the header doctype
const HeaderDoc = `<!DOCTYPE xmeml>` + "\n"

// GenerateXML func will generate an xml
func GenerateXML(jam models.Jam) (*models.XML, error) {
	var xmlProject = models.XML{
		V:       "5",
		Project: makeProject(jam),
	}

	xmlFile, _ := xml.MarshalIndent(xmlProject, "", `   `)

	xmlFile = []byte(Header + HeaderDoc + string(xmlFile))

	err := ioutil.WriteFile("temp/"+jam.Name+".xml", xmlFile, 0644)
	if err != nil {
		return nil, err
	}

	return &xmlProject, nil
}

func makeProject(jam models.Jam) models.Project {
	return models.Project{
		Name:     jam.Name,
		Children: makeChildren(jam),
	}
}

func makeChildren(jam models.Jam) models.Children {
	return models.Children{
		Sequence: makeSequence(jam),
	}
}

func makeSequence(jam models.Jam) models.Sequence {
	return models.Sequence{
		ID:       "sequence-1",
		Duration: int64(calculateDuration(jam.Recordings)),
		Rate:     makeSequeRate(),
		Name:     jam.Name,
		Media:    makeMedia(jam.Recordings),
	}
}

func makeSequeRate() models.Rate {
	return models.Rate{
		TimeBase: 30,
		Ntsc:     false,
	}
}

func makeMedia(rs []models.Recordings) models.Media {
	return models.Media{Audio: makeAudio(rs)}
}

func makeAudio(r []models.Recordings) models.Audio {
	return models.Audio{
		Format:  makeFormat(),
		Outputs: makeOutputs(),
		Tracks:  makeTracks(r),
	}
}
func makeFormat() models.Format {
	return models.Format{
		Samplecharacteristics: models.Samplecharacteristics{
			Depth:      32,
			Samplerate: 12000,
		},
	}
}

func makeOutputs() models.Outputs {
	return models.Outputs{
		Groups: makeGroups(currentJam.Recordings),
	}
}
func makeGroups(rds []models.Recordings) []models.Group {
	var grounds []models.Group
	for i := 0; i < len(rds); i++ {
		group := models.Group{
			Index:       int64(i),
			Downmix:     0,
			Numchannels: 1,
			Channel: models.Channel{
				Index: int64(i),
			},
		}
		grounds = append(grounds, group)
	}
	return grounds
}

func makeTracks(rd []models.Recordings) []models.Track {
	var tracks []models.Track
	for i, r := range rd {
		track := models.Track{
			Enable:   true,
			Locked:   false,
			Clipitem: makeClipitem(r, i),
		}
		tracks = append(tracks, track)
	}
	return tracks
}

func makeClipitem(rd models.Recordings, i int) models.Clipitem {
	return models.Clipitem{
		ID:       "clipitem-" + string(i),
		Name:     rd.User.Name,
		Enabled:  true,
		Duration: int64(convertTime(rd)) * 30,
	}
}
func calculateDuration(r []models.Recordings) float64 {
	return extractLongestDuration(r) * 30
}

func setStartTime(r models.Recordings) {
	//MARK:TODO
}
func setEndTime(r models.Recordings) {

}
func extractLongestDuration(r []models.Recordings) float64 {
	p := fmt.Println
	var l float64
	for _, v := range r {
		if t := convertTime(v); t > l {
			l = t
		}
	}
	p("longest", l)
	return l
}

func isDurationLonger(r models.Recordings) bool {
	for _, cr := range currentJam.Recordings {
		if convertTime(r) > convertTime(cr) {
			return true
		}
	}
	return false
}

func isDurationLess(r models.Recordings) bool {
	for _, cr := range currentJam.Recordings {
		if convertTime(r) < convertTime(cr) {
			return true
		}
	}
	return false
}
func convertTime(r models.Recordings) float64 {
	start, _ := time.Parse("2006-01-02T15:04:05", r.StartTime)
	end, _ := time.Parse("2006-01-02T15:04:05", r.EndTime)

	duration := start.Sub(end).Seconds()
	return math.Abs(duration)
}

func durationInFrames(d time.Duration) float64 {
	return float64(timeToFrame(d))
}
func timeToFrame(t time.Duration) time.Duration {

	return time.Duration(math.Abs(float64(t * 30)))
}
