package models

import "encoding/xml"

// Project struct holds the xml file
type Project struct {
	XMLName  xml.Name `xml:"project"`
	Name     string   `xml:"name"`
	Children Children `xml:"children"`
}
type Children struct {
	Sequence Sequence `xml:"sequence"`
}
type Sequence struct {
	ID       string `xml:"id,attr"`
	Duration int64  `xml:"duration"`
	Rate     Rate   `xml:"rate"`
	Name     string `xml:"name"`
	Media    Media  `xml:"media"`
}
type Rate struct {
	TimeBase int64 `xml:"timebase"`
	Ntsc     bool  `xml:"ntsc"`
}

type Media struct {
	Audio Audio `xml:"audio"`
}
type Audio struct {
	Format  Format  `xml:"format"`
	Outputs Outputs `xml:"outputs"`
	Tracks  []Track `xml:"track"`
}
type Format struct {
	Samplecharacteristics Samplecharacteristics `xml:"samplecharacteristics"`
}
type Samplecharacteristics struct {
	Depth      int64 `xml:"depth"`
	Samplerate int64 `xml:"samplerate"`
}
type Outputs struct {
}
type Track struct {
}
