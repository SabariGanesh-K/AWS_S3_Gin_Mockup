package api
import (
    "mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"net/http"
	db "github.com/SabariGanesh-K/21BPS1209_Backend.git/db/sqlc"


	"fmt"
)
type uploadFileRequest struct {
	File *multipart.FileHeader `form:"file" `
	Filename string `form:"filename" binding:"required"`
}

func (server *Server) uploadFile(ctx *gin.Context) {
	var req uploadFileRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	})
	if err != nil {
		fmt.Printf("Failed to initialize new session: %v", err)
		return
	}

	bucketName := "elasticbeanstalk-us-east-1-686995207617"
	uploader := s3manager.NewUploader(sess)

	file, err := req.File.Open()
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
		return
	}
	defer file.Close()

	_, err = db.UploadFile(uploader, file, bucketName, req.Filename)
	if err != nil {
		fmt.Printf("Failed to upload file: %v", err)
	}

	fmt.Println("Successfully uploaded file!")
}


