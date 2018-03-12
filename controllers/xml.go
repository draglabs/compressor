package controllers

import (
	"compressor/models"
	"encoding/xml"
	"io/ioutil"
	"strconv"
	"time"
	"fmt"
)

// Header is the header of the xml
const Header = `<?xml version="1.0" encoding="UTF-8" standalone="no" ?>` + "\n"

var duration int64         // in frames
var currentEndTime float64 // in frames

// holds the date time pattern
const dateTimePatternOld = "2006-01-02 15:04:05 -0700"
const dateTimePattern = "2006-01-02 15.04.05 -0700"

// HeaderDoc is the header doctype
const HeaderDoc = `<!DOCTYPE xmeml>` + "\n"

// GenerateXML func will generate an xml
//
// Param: jam, a jam model
//
// Returns: a pointer to an xml file model and an error if something went wrong
func GenerateXML(jam models.Jam) (*models.XML, error) {
	fmt.Println("Entering Function GenerateXML")

	var xmlProject = models.XML {
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

// Makes the project section of the xml
//
// Param: jam, a jam model
//
// Returns: a Project model
func makeProject(jam models.Jam) models.Project {
	fmt.Println("Entering Function makeProject")

	return models.Project{
		Name:     jam.Name,
		Children: makeChildren(jam),
	}
}

// Makes the children sequence
//
// Param: jam, a jam model
//
// Returns: a Children model
func makeChildren(jam models.Jam) models.Children {
	fmt.Println("Entering Function makeChildren")

	return models.Children{
		Sequence: makeSequence(jam),
	}
}

// Makes the sequence
//
// Param: jam, a jam model
//
// Returns: a Sequence model
func makeSequence(jam models.Jam) models.Sequence {
	fmt.Println("Entering Function makeSequence")

	duration = globalDuration()
	//jam.Recording = sortByUser(jam.Recording)
	return models.Sequence{
		ID:       "sequence-1",
		Duration: duration,
		Rate:     makeRate(),
		Name:     jam.Name,
		Media:    makeMedia(jam.Recordings),
		TimeCode: makeTimeCode(),
	}
}

// Makes the rate
//
// Param: nothing
//
// Returns: a Rate model; TimeBase = 30, Ntsc = false
func makeRate() models.Rate {
	fmt.Println("Entering Function makeRate")

	return models.Rate{
		TimeBase: 30,
		Ntsc:     false,
	}
}

// Makes the media
//
// Param: an array of Recording models
//
// Returns: a Media model
func makeMedia(recordings []models.Recording) models.Media {
	fmt.Println("Entering Function makeMedia")

	return models.Media{Audio: makeAudio(recordings)}
}

// Makes the time code
//
// Param: nothing
//
// Returns: a TimeCode model; Rate = makeRate(), Frame = 30, Displayformat = NDF
func makeTimeCode() models.TimeCode {
	fmt.Println("Entering Function makeTimeCode")

	return models.TimeCode{
		Rate:          makeRate(),
		Frame:         30,
		Displayformat: "NDF",
	}
}

// Makes the audio
//
// Param: an array of Recording objects
//
// Returns: an Audio model
func makeAudio(recordings []models.Recording) models.Audio {
	fmt.Println("Entering Function makeAudio")

	return models.Audio{
		Format:  makeFormat(),
		Outputs: makeOutputs(),
		Tracks:  makeTracks(recordings),
	}
}

// Makes the format
//
// Param: nothing
//
// Returns: a Format model; Depth = 32, Samplerate = 12000
func makeFormat() models.Format {
	fmt.Println("Entering Function makeFormat")

	return models.Format{
		Samplecharacteristics: models.Samplecharacteristics{
			Depth:      32,
			Samplerate: 12000, // should this be modified and made larger to represent high-quality recordings
		},
	}
}

// Makes the outputs
//
// Param: nothing
//
// Returns: an Outputs model containing the current jam's recordings
func makeOutputs() models.Outputs {
	fmt.Println("Entering Function makeOutputs")

	return models.Outputs{
		Groups: makeGroups(currentJam.Recordings),
	}
}

// Makes the groups
//
// Param: an array of Recording models, unused for some reason
//
// Returns: an array of Group models
func makeGroups(recordings []models.Recording) []models.Group {
	fmt.Println("Entering Function makeGroups")

	var groups []models.Group
	for i := 1; i < 3; i++ {
		group := models.Group {
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

// Makes the tracks
//
// Params: an array of Recording models
//
// Returns: an array of Track models
func makeTracks(recordings []models.Recording) []models.Track {
	fmt.Println("Entering Function makeTracks")

	var tracks []models.Track

	for index, recording := range recordings {
		track := models.Track{
			Enable:             true,
			Locked:             false,
			Clipitem:           makeClipItem(recording, index),
			Outputchannelindex: 25,
		}
		tracks = append(tracks, track)
	}
	return tracks
}

// Makes a ClipItem
//
// Params: recording, a single Recording model; index, the index in the tracks
//
// Returns: a Clipitem model
func makeClipItem(recording models.Recording, index int) models.Clipitem {
	fmt.Println("Entering Function makeClipItem")

	return models.Clipitem{
		ID:           "clipitem-" + strconv.Itoa(index),
		Name:         recording.ID + recording.User.FirstName,
		Enabled:      true,
		Duration:     int64(MakeDuration(recording)),
		Start:        setStartTime(recording),
		End:          setEndTime(recording),
		In:           0,
		Out:          int64(MakeDuration(recording)),
		File:         makeFile(recording, index),
		Sourcetrack:  makeSourceTrack(recording, index),
		Channelcount: int64(index),
	}
}

// Makes the file, appends .caf.
//
// File Name = user's first name + .caf --> Example: "John.caf"
//
// Pathurl = start time + user's first name + .caf --> Example: "20180315 1326 John.caf"
//
// Params: recording, a Recording model; index, the index
//
// Returns: a File model
func makeFile(recording models.Recording, index int) models.File {
	fmt.Println("Entering Function makeFile")

	return models.File {
		ID:       strconv.Itoa(index),
		Name:     recording.User.FirstName + ".caf",
		Pathurl:  recording.StartTime + recording.User.FirstName + ".caf",
		Rate:     makeRate(),
		Duration: int64(MakeDuration(recording)),
		Media:    makeTrackMedia(),
	}
}

// Makes track media
//
// Param: nothing
//
// Returns: a TrackMedia model
func makeTrackMedia() models.TrackMedia {
	fmt.Println("Entering Function makeTrackMedia")

	return models.TrackMedia{Audio: makeTrackAudio()}
}

// Makes track audio
//
// Param: nothing
//
// Returns: a TrackAudio model; Depth = 32, Samplerate = 12000
func makeTrackAudio() models.TrackAudio {
	fmt.Println("Entering Function makeTrackAudio")

	return models.TrackAudio{
		Samplecharacteristics: models.Samplecharacteristics{
			Depth:      32,
			Samplerate: 12000,
		},
	}
}

// Makes the source track
//
// Params: recording, a Recording model (unused); index, the index
//
// Returns: a Sourcetrack model
func makeSourceTrack(recording models.Recording, index int) models.Sourcetrack {
	fmt.Println("Entering Function makeSourceTrack")

	return models.Sourcetrack{
		MediaType:  "audio",
		Trackindex: int64(index),
	}
}

// Sets the recording's start time
//
// Param: a Recording model
//
// Returns: a 64-bit integer with an offset
func setStartTime(recording models.Recording) int64 {
	fmt.Println("Entering Function setStartTime")

	offset := convertStartTime(recording).Sub(zeroStart()).Seconds()
	return int64(offset * 30)
}

// Sets the recording's end time
//
// Param: a Recording model
//
// Returns: a 64-bit integer with an offset
func setEndTime(r models.Recording) int64 {
	fmt.Println("Entering Function setEndTime")

	start := globalDuration()
	d := int64(MakeDuration(r))
	return start + d
}

// MakeDuration calculates the duration and converts it to seconds
//
// Param: a Recording model
//
// Returns: a 64-bit float that represents the recording duration in seconds
func MakeDuration(recording models.Recording) float64 {
	fmt.Println("Entering Function makeDuration")

	//2006-01-02 15:04:05 -0700
	//2006-01-02T15:04:05
	start, _ := time.Parse(dateTimePattern, recording.StartTime)
	end, _ := time.Parse(dateTimePattern, recording.EndTime)
	duration := end.Sub(start)
	return duration.Seconds() * 30
}

// Zeros the start time so that they're all from one spot
//
// Param: nothing
//
// Returns: a Time object with nanosecond precision
func zeroStart() time.Time {
	fmt.Println("Entering Function zeroStart")

	start, _ := time.Parse(dateTimePattern, currentJam.Recordings[0].StartTime)
	return start
}

// Represents the duration of the whole Audition project
//
// Param: nothing
//
// Returns: a 64-bit integer of the total seconds
func globalDuration() int64 {
	fmt.Println("Entering Function globalDuration")

	start, _ := time.Parse(dateTimePattern, currentJam.Recordings[0].StartTime)
	t := currentJam.Recordings[len(currentJam.Recordings)-1]
	end, _ := time.Parse(dateTimePattern, t.EndTime)
	d := end.Sub(start).Seconds()
	return int64(d) * 30
}

// Converts the start time into seconds
//
// Param: a Recording model
//
// Returns: a Time object with nanosecond precision
func convertStartTime(recording models.Recording) time.Time {
	fmt.Println("Entering Function convertStartTime")

	start, _ := time.Parse(dateTimePattern, recording.StartTime)
	return start
}

// sorting by User
// FACT : the Big O is : log(n*n)
// but I don't think will have more than 100
//
// Param: an array of Recording models
//
// Returns: an array of sorted Recording models
func sortByUser(recordings []models.Recording) []models.Recording {
	fmt.Println("Entering Function sortByUser")

	var sorted []models.Recording
	for i := 0; i < len(recordings); i++ {
		x := i
		for j := i; j < len(recordings); j++ {
			if recordings[x].UserID == recordings[j].UserID {
				x = j
			}
		}
		sorted = append(sorted, recordings[x])
	}
	return sorted
}

// Sorting by start time
//
// Param: an array of Recording models
//
// Returns: an array of sorted Recording models
func sortByStartTime(recordings []models.Recording) {
	fmt.Println("Entering Function sortByStartTime")


}
