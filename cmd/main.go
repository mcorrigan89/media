package main

import (
	"os"
	"sync"

	"github.com/mcorrigan89/identity/internal/config"
	"github.com/mcorrigan89/identity/internal/storage"
	"github.com/rs/zerolog"
)

type application struct {
	config  config.Config
	wg      *sync.WaitGroup
	logger  *zerolog.Logger
	storage *storage.StorageService
	// services    *services.Services
	// protoServer *api.ProtoServer
}

func main() {

	logger := getLogger()

	logger.Info().Msg("Starting server")

	cfg := config.Config{}
	config.LoadConfig(&cfg)

	db, err := openDBPool(&cfg)
	if err != nil {
		logger.Err(err).Msg("Failed to open database connection")
		os.Exit(1)
	}
	defer db.Close()

	wg := sync.WaitGroup{}

	// repositories := repositories.NewRepositories(db, &logger, &wg)
	// services := services.NewServices(&repositories, &cfg, &logger, &wg)
	// protoServer := api.NewProtoServer(&cfg, &logger, &wg, &services)

	s := storage.NewStorageService(&logger, cfg.Storage.Endpoint, cfg.Storage.BucketName, cfg.Storage.AccessKeyID, cfg.Storage.SecretAccessKey)

	app := &application{
		wg:      &wg,
		config:  cfg,
		logger:  &logger,
		storage: s,
		// services:    &services,
		// protoServer: protoServer,
	}

	err = app.serve()
	if err != nil {
		logger.Err(err).Msg("Failed to start server")
		os.Exit(1)
	}
}
