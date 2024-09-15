package db

import (
	"os"
	"github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"

)

func UploadFile(uploader *s3manager.Uploader, filePath string, bucketName string, fileName string) (string,error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "",err
	}


	defer file.Close()

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})


	return "url", err
}



