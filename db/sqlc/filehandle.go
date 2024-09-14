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


// func main() {
// 	sess, err := session.NewSessionWithOptions(session.Options{
// 		Profile: "default",
// 		Config: aws.Config{
// 			Region: aws.String("us-east-1"),
// 		},
// 	})

// 	if err != nil {
// 		fmt.Printf("Failed to initialize new session: %v", err)
// 		return
// 	}

// 	bucketName := "elasticbeanstalk-us-east-1-686995207617"
// 	uploader := s3manager.NewUploader(sess)
// 	filename := "1.jpg"

// 	err = UploadFile(uploader, "1.jpg", bucketName, filename)
// 	if err != nil {
// 		fmt.Printf("Failed to upload file: %v", err)
// 	}

// 	fmt.Println("Successfully uploaded file!")
// }

