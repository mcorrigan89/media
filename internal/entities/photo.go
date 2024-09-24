package entities

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mcorrigan89/media/internal/repositories/models"
)

var (
	ErrPhotoNotFound = fmt.Errorf("photo not found")
)

type Photo struct {
	ID      uuid.UUID
	bucket  string
	assetID string
	Width   int32
	Height  int32
	Size    int32
}

func NewPhotoEntityFromModel(model models.Photo) *Photo {
	return &Photo{
		ID:      model.ID,
		bucket:  model.Bucket,
		assetID: model.AssetID,
		Width:   model.Width,
		Height:  model.Height,
		Size:    model.Size,
	}
}

func (p *Photo) UrlSlug() string {
	url := fmt.Sprintf("/%s/%s", p.bucket, p.assetID)
	return url
}
