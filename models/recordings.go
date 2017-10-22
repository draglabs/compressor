package models

type Recordings struct {
	User      Creator `json:"user" bson:"user"`
	FileName  string  `json:"file_name" bson:"file_name"`
	JamID     string  `json:"jam_id" bson:"jamID"`
	StartTime string  `json:"start_time" bson:"startTime"`
	EndTime   string  `json:"end_time" bson:"endTime"`
	Notes     string  `json:"notes" bson:"notes"`
	S3url     string  `json:"s3url" bson:"s3url"`
}
type Creator struct {
	ID    string `json:"id" bson:"id"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"name"`
}
