package services

import (
	"sync"

	"github.com/mcorrigan89/media/internal/config"
	"github.com/mcorrigan89/media/internal/repositories"
	"github.com/rs/zerolog"
)

type ServicesUtils struct {
	logger *zerolog.Logger
	wg     *sync.WaitGroup
	config *config.Config
}

type Services struct {
	utils          *ServicesUtils
	StorageService *StorageService
	PhotoService   *PhotoService
}

func NewServices(repositories *repositories.Repositories, cfg *config.Config, logger *zerolog.Logger, wg *sync.WaitGroup) Services {
	utils := ServicesUtils{
		logger: logger,
		wg:     wg,
		config: cfg,
	}

	storageService := NewStorageService(&utils, repositories.StorageRepository)
	photoService := NewPhotoService(&utils, repositories, storageService)

	return Services{
		utils:          &utils,
		StorageService: storageService,
		PhotoService:   photoService,
	}
}
