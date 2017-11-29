package controllers

import (
	"compressor/models"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

// Header is the header of the xml
const Header = `<?xml version="1.0" encoding="UTF-8" standalone="no" ?>` + "\n"

var duration float64       // in frames
var currentEndTime float64 // in frames

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

	err := ioutil.WriteFile(jam.Name+"/"+jam.Name+".xml", xmlFile, 0644)
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
	duration = calculateDuration(jam.Recordings)
	//jam.Recordings = sortByUser(jam.Recordings)
	return models.Sequence{
		ID:       "sequence-1",
		Duration: int64(duration),
		Rate:     makeRate(),
		Name:     jam.Name,
		Media:    makeMedia(jam.Recordings),
		TimeCode: makeTimeCode(),
	}
}

func makeRate() models.Rate {
	return models.Rate{
		TimeBase: 30,
		Ntsc:     false,
	}
}

func makeMedia(rs []models.Recordings) models.Media {
	return models.Media{Audio: makeAudio(rs)}
}

func makeTimeCode() models.TimeCode {
	return models.TimeCode{
		Rate:          makeRate(),
		Frame:         30,
		Displayformat: "NDF",
	}
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
	for i := 1; i < 3; i++ {
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
	currendID := rd[0].UserID
	for i, r := range rd {

		if currendID != r.UserID {
			currentEndTime = 0
		}
		track := models.Track{
			Enable:             true,
			Locked:             false,
			Clipitem:           makeClipitem(r, i),
			Outputchannelindex: 25,
		}
		tracks = append(tracks, track)
		setCurrentEnd(r)
		currendID = r.UserID
	}
	return tracks
}
func setCurrentEnd(r models.Recordings) {
	if currentEndTime == duration {
		return
	}
	if currentEndTime < duration {
		currentEndTime += MakeDuration(r)
	}
}
func makeClipitem(rd models.Recordings, i int) models.Clipitem {
	return models.Clipitem{
		ID:           "clipitem-" + strconv.Itoa(i),
		Name:         rd.ID + rd.User.FirstName,
		Enabled:      true,
		Duration:     int64(MakeDuration(rd)),
		Start:        int64(currentEndTime),
		End:          int64(currentEndTime + setEndTime(rd)),
		In:           0,
		Out:          int64(MakeDuration(rd)),
		File:         makeFile(rd, i),
		Sourcetrack:  makeSourceTrack(rd, i),
		Channelcount: int64(i),
	}
}
func makeFile(r models.Recordings, i int) models.File {
	return models.File{
		ID:       strconv.Itoa(i),
		Name:     r.User.FirstName + ".caf",
		Pathurl:  r.StartTime + r.User.FirstName + ".caf",
		Rate:     makeRate(),
		Duration: int64((MakeDuration(r))),
		Media:    makeTrackMedia(),
	}
}

func makeTrackMedia() models.TrackMedia {
	return models.TrackMedia{Audio: makeTrackAudio()}
}
func makeTrackAudio() models.TrackAudio {
	return models.TrackAudio{
		Samplecharacteristics: models.Samplecharacteristics{
			Depth:      32,
			Samplerate: 12000,
		},
	}
}
func makeSourceTrack(r models.Recordings, i int) models.Sourcetrack {
	return models.Sourcetrack{
		MediaType:  "audio",
		Trackindex: int64(i),
	}
}

// calculateDuration func, should returns the logest
// duration in frames
func calculateDuration(r []models.Recordings) float64 {
	var l float64
	for _, v := range r {
		t := MakeDuration(v)
		l += t
	}
	fmt.Println("Current Duration", l)
	duration = l
	return l * 30
}

func setStartTime(r models.Recordings) float64 {
	offset := duration - MakeDuration(r)
	return offset
}
func setEndTime(r models.Recordings) float64 {
	return MakeDuration(r)
}

func isDurationLonger(r models.Recordings) bool {
	for _, cr := range currentJam.Recordings {
		if MakeDuration(r) > MakeDuration(cr) {
			return true
		}
	}
	return false
}

func isDurationLess(r models.Recordings) bool {
	for _, cr := range currentJam.Recordings {
		if MakeDuration(r) < MakeDuration(cr) {
			return true
		}
	}
	return false
}

// MakeDuration calculates the duration and
// convert is to seconds
func MakeDuration(r models.Recordings) float64 {
	//2006-01-02 15:04:05 -0700
	//2006-01-02T15:04:05
	start, _ := time.Parse("2006-01-02 15:04:05 -0700", r.StartTime)
	end, _ := time.Parse("2006-01-02 15:04:05 -0700", r.EndTime)
	duration := end.Sub(start)
	return duration.Seconds() * 30
}

// sorting by User
//FUN FACT : the Big O is : log(n*n)
// but i dont think will have more than 100
func sortByUser(rs []models.Recordings) []models.Recordings {
	var sorted []models.Recordings
	for i := 0; i < len(rs); i++ {
		x := i
		for j := i; j < len(rs); j++ {
			if rs[x].UserID == rs[j].UserID {
				x = j
			}
		}
		sorted = append(sorted, rs[x])
	}
	return sorted
}
func sortByStartTime(rs []models.Recordings) {

}
