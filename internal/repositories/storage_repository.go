package repositories

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type StorageRepository struct {
	utils *ServicesUtils
}

func NewStorageRepository(utils *ServicesUtils) *StorageRepository {
	return &StorageRepository{
		utils: utils,
	}
}

func (s *StorageRepository) GetObject(ctx context.Context, objectKey string) ([]byte, error) {

	var useSSL = true

	minioClient, err := minio.New(s.utils.config.Storage.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(s.utils.config.Storage.AccessKeyID, s.utils.config.Storage.SecretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		s.utils.logger.Err(err).Ctx(ctx).Msg("Failed to create minio client")
		return nil, err
	}

	objResponse, err := minioClient.GetObject(ctx, s.utils.config.Storage.BucketName, objectKey, minio.GetObjectOptions{})
	if err != nil {
		s.utils.logger.Err(err).Ctx(ctx).Msg("Failed to get object from storage")
		return nil, err
	}

	defer objResponse.Close()

	respByte, err := io.ReadAll(objResponse)
	if err != nil {
		s.utils.logger.Err(err).Msg("Failed to read object from storage")
		return nil, err
	}

	return respByte, nil
}
