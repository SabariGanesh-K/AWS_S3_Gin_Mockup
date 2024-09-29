package db

import (
	"github.com/aws/aws-sdk-go/aws"
//  "github.com/aws/aws-sdk-go/aws/session"
"fmt"
"io"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

)

	
func UploadFile(uploader *s3manager.Uploader,  file io.Reader, bucketName string,fileName string) (string, error) {


   

	_, err := uploader.Upload(&s3manager.UploadInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(fileName),
        Body:   file,
    })

    if err != nil {
        return "", fmt.Errorf("error uploading file to S3: %w", err)
    }

    return "", nil
}




