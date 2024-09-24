package repositories

import (
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mcorrigan89/media/internal/config"
	"github.com/mcorrigan89/media/internal/repositories/models"

	"github.com/rs/zerolog"
)

const defaultTimeout = 10 * time.Second

type ServicesUtils struct {
	logger *zerolog.Logger
	config *config.Config
	wg     *sync.WaitGroup
	db     *pgxpool.Pool
}

type Repositories struct {
	utils             ServicesUtils
	StorageRepository *StorageRepository
	PhotoRepository   *PhotoRepository
}

func NewRepositories(db *pgxpool.Pool, cfg *config.Config, logger *zerolog.Logger, wg *sync.WaitGroup) Repositories {

	queries := models.New(db)
	utils := ServicesUtils{
		logger: logger,
		config: cfg,
		wg:     wg,
		db:     db,
	}

	storageRepo := NewStorageRepository(&utils)
	photoRepo := NewPhotoRepository(&utils, db, queries)

	return Repositories{
		utils:             utils,
		StorageRepository: storageRepo,
		PhotoRepository:   photoRepo,
	}
}
