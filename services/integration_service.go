package services

import (
	"bytes"
	"context"
	"fmt"
	"gohighload/models"
	"gohighload/utils"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func init() {
	var err error
	minioClient, err = minio.New("localhost:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("minioadmin", "minioadmin", ""),
		Secure: false,
	})
	if err != nil {
		utils.Logger.Printf("Failed to initialize MinIO client: %v\n", err)
		return
	}
	ctx := context.Background()
	exists, errBucket := minioClient.BucketExists(ctx, "notifications")
	if errBucket != nil {
		utils.Logger.Printf("Error checking bucket: %v\n", errBucket)
		return
	}
	if !exists {
		err = minioClient.MakeBucket(ctx, "notifications", minio.MakeBucketOptions{})
		if err != nil {
			utils.Logger.Printf("Failed to create bucket: %v\n", err)
		} else {
			utils.Logger.Printf("Created bucket 'notifications'\n")
		}
	}
}

func LogUserAction(action string, userID int) {
	go func() {
		utils.Logger.Printf("Audit Log: %s for user ID %d\n", action, userID)
	}()
}

func SendNotification(user models.User, action string) {
	go func() {
		message := fmt.Sprintf("Notification for %s: %s\n", action, user.Email)
		bucketName := "notifications"
		objectName := fmt.Sprintf("notification-%d-%s.txt", user.ID, action)
		reader := bytes.NewReader([]byte(message))
		_, err := minioClient.PutObject(context.Background(), bucketName, objectName, reader, int64(len(message)), minio.PutObjectOptions{ContentType: "text/plain"})
		if err != nil {
			utils.Logger.Printf("Failed to upload to MinIO: %v\n", err)
		} else {
			utils.Logger.Printf("Uploaded notification to MinIO: %s\n", objectName)
		}
	}()
}
