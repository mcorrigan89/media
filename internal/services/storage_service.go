package services

import (
	"context"
	"io"

	"github.com/mcorrigan89/media/internal/repositories"
)

type StorageService struct {
	utils             *ServicesUtils
	storageRepository *repositories.StorageRepository
}

func NewStorageService(utils *ServicesUtils, storageRepository *repositories.StorageRepository) *StorageService {
	return &StorageService{
		utils:             utils,
		storageRepository: storageRepository,
	}
}

func (s *StorageService) GetObject(ctx context.Context, objectKey string) ([]byte, error) {
	s.utils.logger.Info().Ctx(ctx).Str("objectKey", objectKey).Msg("Getting object from storage")

	imageBytes, err := s.storageRepository.GetObject(ctx, objectKey)
	if err != nil {
		s.utils.logger.Err(err).Ctx(ctx).Msg("Failed to get object from storage")
		return nil, err
	}

	return imageBytes, nil
}

func (s *StorageService) UploadObject(ctx context.Context, objectKey string, object io.Reader, size int64) (*string, error) {
	s.utils.logger.Info().Ctx(ctx).Str("objectKey", objectKey).Msg("Uploading object to storage")

	assetId, err := s.storageRepository.UploadObject(ctx, objectKey, object, size)
	if err != nil {
		s.utils.logger.Err(err).Ctx(ctx).Msg("Failed to upload object to storage")
		return nil, err
	}

	return assetId, nil
}
