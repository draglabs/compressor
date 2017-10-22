package models

import "encoding/xml"

// Project struct holds the xml file
type Project struct {
	XMLName  xml.Name `xml:"project"`
	Name     string   `xml:"name"`
	Children Children `xml:"children"`
}

//Children struct
type Children struct {
	Sequence Sequence `xml:"sequence"`
}

//Sequence struct, represent the sequence for the project
type Sequence struct {
	ID       string   `xml:"id,attr"`
	Duration int64    `xml:"duration"`
	Rate     Rate     `xml:"rate"`
	Name     string   `xml:"name"`
	Media    Media    `xml:"media"`
	TimeCode TimeCode `xml:"timecode"`
}

// Rate struct, holds the rate
type Rate struct {
	TimeBase int64 `xml:"timebase"`
	Ntsc     bool  `xml:"ntsc"`
}

// Media struct, holds the audio
type Media struct {
	Audio Audio `xml:"audio"`
}

// Audio struct, holds format and the traks
type Audio struct {
	Format  Format  `xml:"format"`
	Outputs Outputs `xml:"outputs"`
	Tracks  []Track `xml:"track,omitempty"`
}

// Format struct
type Format struct {
	Samplecharacteristics Samplecharacteristics `xml:"samplecharacteristics"`
}

// Samplecharacteristics struct
type Samplecharacteristics struct {
	Depth      int64 `xml:"depth"`
	Samplerate int64 `xml:"samplerate"`
}

//Outputs struct
type Outputs struct {
	Groups []Group `xml:"group"`
}
type Group struct {
	Index       int64   `xml:"index"`
	Numchannels int64   `xml:"numchannels"`
	Downmix     int64   `xml:"downmix"`
	Channel     Channel `xml:"channel"`
}

type Channel struct {
	Index int64 `xml:"index"`
}

//Track struct, represent the tracks on the project
type Track struct {
	Enable             bool     `xml:"enable"`
	Locked             bool     `xml:"locked"`
	Clipitem           Clipitem `xml:"clipitem"`
	Outputchannelindex int      `xml:"outputchannelindex"`
}
type Clipitem struct {
	ID           string      `xml:"id,attr"`
	Name         string      `xml:"name"`
	Enabled      bool        `xml:"enabled"`
	Duration     int64       `xml:"duration"`
	Start        int64       `xml:"start"`
	End          int64       `xml:"end"`
	In           int64       `xml:"in"`
	Out          int64       `xml:"out"`
	File         File        `xml:"file"`
	Sourcetrack  Sourcetrack `xml:"sourcetrack"`
	Channelcount int64       `xml:"channelcount"`
}
type File struct {
	ID       string     `xml:"id,attr"`
	Name     string     `xml:"name"`
	Pathurl  string     `xml:"pathurl"`
	Rate     Rate       `xml:"rate"`
	Duration int64      `xml:"duration"`
	Media    TrackMedia `xml:"media"`
}
type TrackMedia struct {
	Audio TrackAudio `xml:"audio"`
}
type TrackAudio struct {
	Samplecharacteristics Samplecharacteristics `xml:"samplecharacteristics"`
}
type Sourcetrack struct {
}

//TimeCode struct
type TimeCode struct {
}
