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
	var project = models.XML{
		V:       "5",
		Project: makeProject(jam),
	}
	xmlFile, _ := xml.MarshalIndent(project, "", `   `)
	xmlFile = []byte(Header + HeaderDoc + string(xmlFile))
	err := ioutil.WriteFile("temp/"+jam.Name+".xml", xmlFile, 0644)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%s", xmlFile)
	return &project, nil
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
		Name:     jam.Name,
		Duration: int64(calculateDuration(jam.Recordings)),
		TimeCode: models.TimeCode{
			Rate: models.Rate{
				TimeBase: 30,
				Ntsc:     false,
			},
			Frame:         0,
			Displayformat: "NDF",
		},
	}
}
func calculateDuration(r []models.Recordings) float64 {
	return extractLongestDuration(r) * 30
}

func extractLongestDuration(r []models.Recordings) float64 {
	p := fmt.Println
	var l float64
	for _, v := range r {
		if t := convertTime(v); t > l {
			l = t
		}
	}
	p("logest", l)
	return l
}
func convertTime(r models.Recordings) float64 {
	start, _ := time.Parse("2006-01-02T15:04:05", r.StartTime)
	end, _ := time.Parse("2006-01-02T15:04:05", r.EndTime)

	duration := start.Sub(end).Seconds()
	return math.Abs(duration)

}
