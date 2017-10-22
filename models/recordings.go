package models

type Recordings struct {
	User      Creator `json:"user" bson:"user"`
	FileName  string  `json:"file_name" bson:"file_name"`
	JamID     string
	StartTime string
	EndTime   string
	Notes     string
	S3url     string
}
type Creator struct {
	ID    string
	Name  string
	Email string
}

type UserFB struct {
	facebookID string
	Name       string
	FirstName  string
	LastName   string
	Token      string
	Email      string
}
