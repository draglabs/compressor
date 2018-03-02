package models

// Recording struct
type Recording struct {
	ID        string `json:"id" bson:"_id"`
	UserID    string `json:"user_id" bson:"user_id"`
	FileName  string `json:"file_name"   bson:"file_name"`
	JamID     string `json:"jam_id"      bson:"jam_id"`
	StartTime string `json:"start_time"  bson:"start_time"`
	EndTime   string `json:"end_time"    bson:"end_time"`
	Notes     string `json:"notes"       bson:"notes"`
	S3url     string `json:"s3url"       bson:"s3url"`
	User      User
}
