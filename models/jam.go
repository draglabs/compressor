package models

// Jam struct, models our Jam collection
type Jam struct {
	Pin    string `json:"pin"    bson:"pin"`
	Status bool   `json:"status" bson:"status"`
	Name   string `json:"name"   bson:"name"`

	CollaboartorID []string     `json:"collaborator_id" bson:"collaboartorID"`
	Coordinates    []float64    `json:"coordinates"     bson:"coordinates"`
	Collaborators  []Creator    `json:"collaborators"   bson:"collaborators"`
	Recordings     []Recordings `json:"recordings"      bson:"recordings"`
	Location       string       `json:"location"        bson:"location"`
	Creator        Creator      `json:"creator"         bson:"creator"`
	StartTime      string       `json:"start_time"      bson:"startTime"`
	EndTime        string       `json:"end_time"        bson:"endTime"`
	StatusID       string       `json:"status_id"       bson:"statusID"`
	Notes          string       `json:"notes"           bson:"notes"`
}
