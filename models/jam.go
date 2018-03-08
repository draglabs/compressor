package models

// Jam struct, models our Jam collection
type Jam struct {
	ID            string       `json:"id"             bson:"_id"`
	Pin           string       `json:"pin"             bson:"pin"`
	IsCurrent     bool         `json:"is_current"      bson:"is_current"`
	Name          string       `json:"name"            bson:"name"`
	UserID        string       `json:"user_id"         bson:"user_id"`
	Coordinates   []float64    `json:"coordinates"     bson:"coordinates"`
	Collaborators []User       `json:"collaborators"   bson:"collaborators"`
	Recordings    []Recordings `json:"recordings"      bson:"recordings"`
	Location      string       `json:"location"        bson:"location"`
	StartTime     string       `json:"start_time"      bson:"start_time"`
	EndTime       string       `json:"end_time"        bson:"end_time"`
	Notes         string       `json:"notes"           bson:"notes"`
	Creator       User
	Link          string `json:"link"  bson:"link"`
}
