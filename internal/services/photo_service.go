package services

import (
	"context"
	"io"
	"math"

	"github.com/google/uuid"
	"github.com/h2non/bimg"

	"github.com/mcorrigan89/media/internal/entities"
	"github.com/mcorrigan89/media/internal/repositories"
)

var renditionSizes = map[string]int{
	"small":  320,
	"medium": 640,
	"large":  1024,
	"xlarge": 2048,
}

type PhotoService struct {
	utils           *ServicesUtils
	photoRepository *repositories.PhotoRepository
	storageService  *StorageService
}

func NewPhotoService(utils *ServicesUtils, repositores *repositories.Repositories, storageService *StorageService) *PhotoService {
	return &PhotoService{
		utils:           utils,
		photoRepository: repositores.PhotoRepository,
		storageService:  storageService,
	}
}

func (s *PhotoService) ProcessImage(ctx context.Context, imageBytes []byte, renditionSize string) ([]byte, string, error) {
	s.utils.logger.Info().Ctx(ctx).Msg("Processing Image")

	var width, height int

	img := bimg.NewImage(imageBytes)

	metadata, err := img.Metadata()
	if err != nil {
		s.utils.logger.Err(err).Msg("Failed to get metadata")
		return nil, "", err
	}

	width, height = s.calculateDimensions(metadata.Size.Width, metadata.Size.Height, renditionSizes[renditionSize])

	contentType := "image/webp"
	processedImage, err := img.Process(bimg.Options{
		Height: height,
		Width:  width,
		Type:   bimg.WEBP,
	})
	if err != nil {
		s.utils.logger.Err(err).Ctx(ctx).Msg("Failed to process image")
		return nil, "", err
	}

	return processedImage, contentType, nil
}

func (s *PhotoService) calculateDimensions(width, height, maxSize int) (int, int) {
	aspectRatio := float64(width) / float64(height)

	var newWidth, newHeight int

	if width > height {
		newWidth = maxSize
		newHeight = int(math.Round(float64(newWidth) / aspectRatio))
	} else {
		newHeight = maxSize
		newWidth = int(math.Round(float64(newHeight) * aspectRatio))
	}

	return newWidth, newHeight
}

func (s *PhotoService) imageMetadata(ctx context.Context, imageBytes []byte) (*bimg.ImageMetadata, error) {
	img := bimg.NewImage(imageBytes)

	metadata, err := img.Metadata()
	if err != nil {
		s.utils.logger.Err(err).Ctx(ctx).Msg("Failed to get metadata")
		return nil, err
	}

	return &metadata, nil
}

func (s *PhotoService) GetPhotoByID(ctx context.Context, id uuid.UUID) (*entities.Photo, error) {
	s.utils.logger.Info().Ctx(ctx).Msg("Get photo by ID")

	photo, err := s.photoRepository.GetPhotoByID(ctx, id)
	if err != nil {
		s.utils.logger.Err(err).Msg("Failed to get photo by ID")
		return nil, err
	}

	return photo, nil
}

type CreatePhotoArgs struct {
	Filename string
	File     io.Reader
	Size     int64
}

func (s *PhotoService) CreatePhoto(ctx context.Context, args CreatePhotoArgs) (*entities.Photo, error) {
	s.utils.logger.Info().Ctx(ctx).Msg("Create photo")

	assetId, err := s.storageService.UploadObject(ctx, args.Filename, args.File, args.Size)
	if err != nil {
		s.utils.logger.Err(err).Msg("Failed to upload object to storage")
		return nil, err
	}

	imageBytes, err := io.ReadAll(args.File)
	if err != nil {
		s.utils.logger.Err(err).Ctx(ctx).Msg("Failed to read image bytes")
		return nil, err
	}

	metadata, err := s.imageMetadata(ctx, imageBytes)
	if err != nil {
		s.utils.logger.Err(err).Ctx(ctx).Msg("Failed to get image metadata")
		return nil, err
	}

	photo, err := s.photoRepository.CreatePhoto(ctx, repositories.CreatePhotoArgs{
		Bucket:  "images",
		AssetID: *assetId,
		Width:   int32(metadata.Size.Width),
		Height:  int32(metadata.Size.Height),
		Size:    int32(args.Size),
	})
	if err != nil {
		s.utils.logger.Err(err).Msg("Failed to create photo")
		return nil, err
	}

	return photo, nil
}
