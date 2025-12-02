package services

import (
	"bytes"
	"context"
	"fmt"
	"gohighload/models"

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
		fmt.Printf("Failed to initialize MinIO client: %v\n", err)
	}
}

func LogUserAction(action string, userID int) {
	go func() {
		fmt.Printf("Audit Log: %s for user ID %d\n", action, userID)
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
			fmt.Printf("Failed to upload to MinIO: %v\n", err)
		} else {
			fmt.Printf("Uploaded notification to MinIO: %s\n", objectName)
		}
	}()
}
