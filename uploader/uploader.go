package uploader

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	s3region = "us-west-1"
)

// Upload func, uploads are archive to s3
// once ge a response from the server then
// it moves to call the mailer to send email to the user
func Upload(filename string, id string) (string, error) {

	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession())

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("dsoundboy-export"),
		Key:    aws.String(id),
		Body:   f,
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	return aws.StringValue(&result.Location), nil
}
