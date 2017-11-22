package models

//User struct
type User struct {
	ID        string `json:"id" bson:"_id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	FBEmail   string `json:"fb_email" bson:"fb_email"`
	FBID      string `json:"fb_id" bson:"fb_id"`
	Email     string `json:"email,omitempty" bson:"email,omitempty"`
	// CurrentJam *Jam   `json:"current_jam" bson:"current_jam"`
}

// JamDetails struct
type JamDetails struct {
	ID        string `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	StartTime string `json:"start_time" bson:"startTime"`
	EndTime   string `json:"end_time" bson:"endTime"`
}
