package repositories

import (
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mcorrigan89/media/internal/config"

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
}

func NewRepositories(db *pgxpool.Pool, cfg *config.Config, logger *zerolog.Logger, wg *sync.WaitGroup) Repositories {

	utils := ServicesUtils{
		logger: logger,
		config: cfg,
		wg:     wg,
		db:     db,
	}

	storageRepo := NewStorageRepository(&utils)

	return Repositories{
		utils:             utils,
		StorageRepository: storageRepo,
	}
}
