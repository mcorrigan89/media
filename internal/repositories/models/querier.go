// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package models

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreatePhoto(ctx context.Context, arg CreatePhotoParams) (Photo, error)
	GetPhotoByID(ctx context.Context, id uuid.UUID) (GetPhotoByIDRow, error)
}

var _ Querier = (*Queries)(nil)
