package services

import (
	"context"
	"math"

	"github.com/h2non/bimg"
)

var renditionSizes = map[string]int{
	"small":  320,
	"medium": 640,
	"large":  1024,
	"xlarge": 2048,
}

type PhotoService struct {
	utils *ServicesUtils
}

func NewPhotoService(utils *ServicesUtils) *PhotoService {
	return &PhotoService{
		utils: utils,
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
		s.utils.logger.Err(err).Msg("Failed to process image")
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
