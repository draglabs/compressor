package controllers

import (
	"compressor/models"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// Header is the header of the xml
const Header = `<?xml version="1.0" encoding="UTF-8" standalone="no" ?>` + "\n"

// HeaderDoc is the header doctype
const HeaderDoc = `<!DOCTYPE xmeml>` + "\n"

// GenerateXML func will generate an xml
func GenerateXML() (*models.XML, error) {
	var project = models.XML{
		V:       "5",
		Project: makeProject(),
	}
	xmlFile, _ := xml.MarshalIndent(project, "", `   `)
	xmlFile = []byte(Header + HeaderDoc + string(xmlFile))
	err := ioutil.WriteFile("project.xml", xmlFile, 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", xmlFile)
	return &project, nil
}
func makeProject() models.Project {
	return models.Project{
		Name:     "Jam",
		Children: makeChildren(),
	}
}
func makeChildren() models.Children {
	return models.Children{
		Sequence: makeSequence(),
	}
}
func makeSequence() models.Sequence {
	return models.Sequence{
		ID: "sequence-1",
	}
}
