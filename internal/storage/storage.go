package storage

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type StorageService struct {
	Endpoint        string
	BucketName      string
	AccessKeyID     string
	SecretAccessKey string
}

func (s *StorageService) GetObject(objectKey string) (*minio.Object, error) {
	// var endpoint = "storage.corrigan.io"
	// var bucketName = "images"
	// var accessKeyID = "iJJO2pvme2ECtOIEdDI2"
	// var secretAccessKey = "oDpXWK0xSQ1usMAIhdeOiAF3RyrHyZFXsv9MklSI"
	var useSSL = true

	minioClient, err := minio.New(s.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(s.AccessKeyID, s.SecretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	minioClient.GetObject(context.Background(), s.BucketName, objectKey, minio.GetObjectOptions{})

	obj, err := minioClient.GetObject(context.Background(), s.BucketName, objectKey, minio.GetObjectOptions{})

	if err != nil {
		return nil, err
	}
	return obj, nil
}
