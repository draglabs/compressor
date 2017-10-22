package models

//User struct
type User struct {
	ActiveJam  interface{} `json:"-" bson:"activeJam"`
	ID         string      `json:"id" bson:"_id"`
	Facebook   UserFB      `json:"facebook" bson:"facebook"`
	Recordings Recordings  `json:"recordings" bson:"recordings"`
	JamsIDs    []string    `json:"jams_ids" bson:"jamsIds"`
	JamDetails JamDetails  `json:"jam_details" bson:"jamDetails"`
}

// JamDetails struct
type JamDetails struct {
	ID        string `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	StartTime string `json:"start_time" bson:"startTime"`
	EndTime   string `json:"end_time" bson:"endTime"`
}

// UserFB struct
type UserFB struct {
	FacebookID string `json:"facebook_id" bson:"facebook_id"`
	Name       string `json:"name" bson:"name"`
	FirstName  string `json:"first_name" bson:"first_name"`
	LastName   string `json:"last_name" bson:"last_name"`
	Token      string `json:"-" bson:"token"`
	Email      string `json:"email" bson:"email"`
}
