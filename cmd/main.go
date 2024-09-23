package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/h2non/bimg"
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

	buffer, err := bimg.Read("image.jpeg")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	newImage, err := bimg.NewImage(buffer).Resize(800, 600)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	size, err := bimg.NewImage(newImage).Size()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if size.Width == 800 && size.Height == 600 {
		fmt.Println("The image size is valid")
	}

	bimg.Write("new.jpeg", newImage)

	wg := sync.WaitGroup{}

	// repositories := repositories.NewRepositories(db, &logger, &wg)
	// services := services.NewServices(&repositories, &cfg, &logger, &wg)
	// protoServer := api.NewProtoServer(&cfg, &logger, &wg, &services)

	s := storage.StorageService{
		Endpoint:        cfg.Storage.Endpoint,
		BucketName:      cfg.Storage.BucketName,
		AccessKeyID:     cfg.Storage.AccessKeyID,
		SecretAccessKey: cfg.Storage.SecretAccessKey,
	}

	app := &application{
		wg:      &wg,
		config:  cfg,
		logger:  &logger,
		storage: &s,
		// services:    &services,
		// protoServer: protoServer,
	}

	err = app.serve()
	if err != nil {
		logger.Err(err).Msg("Failed to start server")
		os.Exit(1)
	}
}
