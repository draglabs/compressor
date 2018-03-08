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
func Upload(filename string, key string) (string, error) {

	sess, err := session.NewSession(&aws.Config{Region: aws.String(s3region)})
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
		Key:    aws.String(key + ".zip"),
		Body:   f,
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	return aws.StringValue(&result.Location), nil
}

// CleanupAfterUpload func, will clean up
// the temp dirs and files created during
// the downlaods and archive
func CleanupAfterUpload(temp, archive string) {
	err := os.RemoveAll(temp)
	if err != nil {
		fmt.Println("error deleting temp folder ", err)
	}
	err = os.Remove(archive)
	if err != nil {
		fmt.Println("error removing archive file ", err)
	}

}
