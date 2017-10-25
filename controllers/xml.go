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

var logestDuration float64

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
	var groups []models.Group
	for i := 0; i < len(rds); i++ {
		group := models.Group{
			Index:       int64(i),
			Downmix:     0,
			Numchannels: 1,
			Channel: models.Channel{
				Index: int64(i),
			},
		}
		groups = append(groups, group)
	}
	return groups
}

func makeTracks(rd []models.Recordings) []models.Track {
	var tracks []models.Track
	for i, r := range rd {
		track := models.Track{
			Enable:             true,
			Locked:             false,
			Clipitem:           makeClipitem(r, i),
			Outputchannelindex: 25,
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
		Start:    int64(setStartTime(rd)) * 30,
		End:      int64(setEndTime(rd)) * 30,
	}
}

// calculateDuration func, should returns the logest
// duration in frames
func calculateDuration(r []models.Recordings) float64 {
	return extractLongestDuration(r) * 30
}

func setStartTime(r models.Recordings) float64 {
	offset := logestDuration - convertTime(r)
	return offset
	//MARK:TODO
}
func setEndTime(r models.Recordings) float64 {
	return convertTime(r)
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
	logestDuration = l
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

func durationInFrames(d float64) float64 {
	return float64(timeToFrame(d))
}
func timeToFrame(t float64) float64 {

	return (math.Abs(float64(t * 30)))
}

// sorting by duration
//FUN FACT : the Big O is : log(n*n)
// but i dont think will have more than 100
func sortByLongDuration(rs []models.Recordings) []models.Recordings {
	var sorted []models.Recordings
	for i := 0; i < len(rs); i++ {
		x := i
		for j := i; j < len(rs); j++ {
			if convertTime(rs[x]) > convertTime(rs[j]) {
				x = j
			}
		}

		sorted = append(sorted, rs[x])
	}
	return sorted
}
