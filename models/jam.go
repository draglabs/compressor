package models

// Jam struct, models our Jam collection
type Jam struct {
	Pin    string `json:"pin" bson:"pin"`
	Status bool   `json:"status" bson:"status"`
	Name   string `json:"name" bson:"name"`

	Coordinates    []float64 `json:"coordinates" bson:"coordinates"`
	Location       string    `json:"location" bson:"location"`
	Creator        Creator   `json:"creator" bson:"creator"`
	CollaboartorID []string  `json:"collaborator_id" bson:"collaboratorID"`
	Collaborators  []string
	Recordings     []Recordings
	StartTime      string
	EndTime        string
	StatusID       string
	Notes          string
}
