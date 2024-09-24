package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mcorrigan89/media/internal/entities"
	"github.com/mcorrigan89/media/internal/repositories/models"
)

type PhotoRepository struct {
	utils   *ServicesUtils
	DB      *pgxpool.Pool
	queries *models.Queries
}

func NewPhotoRepository(utils *ServicesUtils, db *pgxpool.Pool, queries *models.Queries) *PhotoRepository {
	return &PhotoRepository{
		utils:   utils,
		DB:      db,
		queries: queries,
	}
}

func (repo *PhotoRepository) GetPhotoByID(ctx context.Context, id uuid.UUID) (*entities.Photo, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	row, err := repo.queries.GetPhotoByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entities.ErrPhotoNotFound
		} else {
			repo.utils.logger.Err(err).Ctx(ctx).Msg("Get photo by ID")
			return nil, err
		}
	}

	entity := entities.NewPhotoEntityFromModel(row.Photo)

	return entity, nil
}

type CreatePhotoArgs struct {
	Bucket  string
	AssetID string
	Width   int32
	Height  int32
	Size    int32
	OwnerID *uuid.UUID
}

func (repo *PhotoRepository) CreatePhoto(ctx context.Context, args CreatePhotoArgs) (*entities.Photo, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	row, err := repo.queries.CreatePhoto(ctx, models.CreatePhotoParams{
		Bucket:  args.Bucket,
		AssetID: args.AssetID,
		Width:   args.Width,
		Height:  args.Height,
		Size:    args.Size,
		OwnerID: args.OwnerID,
	})
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Create photo")
		return nil, err
	}

	entity := entities.NewPhotoEntityFromModel(row)

	return entity, nil
}
