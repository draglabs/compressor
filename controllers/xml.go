package controllers

import (
	"github.com/draglabs/compressor/models"

	"encoding/xml"
	"io/ioutil"
	"strconv"
	"time"
)

// Header is the header of the xml
const Header = `<?xml version="1.0" encoding="UTF-8" standalone="no" ?>` + "\n"

var duration int64         // in frames
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
	duration = globalDuration()
	//jam.Recordings = sortByUser(jam.Recordings)
	return models.Sequence{
		ID:       "sequence-1",
		Duration: duration,
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
		ID:           "clipitem-" + strconv.Itoa(i),
		Name:         rd.ID + rd.User.FirstName,
		Enabled:      true,
		Duration:     int64(MakeDuration(rd)),
		Start:        setStartTime(rd),
		End:          setEndTime(rd),
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

func setStartTime(r models.Recordings) int64 {
	offset := convertStartTime(r).Sub(zeroStart()).Seconds()
	return int64(offset * 30)
}

func setEndTime(r models.Recordings) int64 {
	start := globalDuration()
	d := int64(MakeDuration(r))
	return (start + d)
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
func zeroStart() time.Time {
	start, _ := time.Parse("2006-01-02 15:04:05 -0700", currentJam.Recordings[0].StartTime)
	return start
}
func globalDuration() int64 {
	start, _ := time.Parse("2006-01-02 15:04:05 -0700", currentJam.Recordings[0].StartTime)
	t := currentJam.Recordings[len(currentJam.Recordings)-1]
	end, _ := time.Parse("2006-01-02 15:04:05 -0700", t.EndTime)
	d := end.Sub(start).Seconds()
	return int64(d) * 30
}
func convertStartTime(r models.Recordings) time.Time {
	start, _ := time.Parse("2006-01-02 15:04:05 -0700", r.StartTime)
	return start
}

// sorting by User
//FACT : the Big O is : log(n*n)
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
