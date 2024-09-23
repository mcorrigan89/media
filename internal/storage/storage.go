package storage

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog"
)

type StorageService struct {
	logger          *zerolog.Logger
	Endpoint        string
	BucketName      string
	AccessKeyID     string
	SecretAccessKey string
}

func NewStorageService(logger *zerolog.Logger, endpoint, bucketName, accessKeyID, secretAccessKey string) *StorageService {
	return &StorageService{
		logger:          logger,
		Endpoint:        endpoint,
		BucketName:      bucketName,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
	}
}

func (s *StorageService) GetObject(ctx context.Context, objectKey string) (*minio.Object, error) {

	var useSSL = true

	minioClient, err := minio.New(s.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(s.AccessKeyID, s.SecretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		s.logger.Err(err).Ctx(ctx).Msg("Failed to create minio client")
		return nil, err
	}

	obj, err := minioClient.GetObject(ctx, s.BucketName, objectKey, minio.GetObjectOptions{})
	if err != nil {
		s.logger.Err(err).Ctx(ctx).Msg("Failed to get object from storage")
		return nil, err
	}
	return obj, nil
}
