package main

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess, err := session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
			Profile:           "sc-development",
			Config: aws.Config{
				Region: aws.String("ap-southeast-2"),
			},
		},
	)

	// Create S3 service client
	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("sc-dev-issues-blobqueen"),
		Key:    aws.String("000acb73-08ab-4937-8c08-eef92d1fb6d0"),
	})
	urlStr, err := req.Presign(15 * time.Minute)

	if err != nil {
		log.Println("Failed to sign request", err)
		return
	}

	log.Println("The URL is", urlStr)
}
