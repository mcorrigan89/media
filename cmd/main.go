package main

import (
	"os"
	"sync"

	"github.com/mcorrigan89/media/internal/api"
	"github.com/mcorrigan89/media/internal/config"
	"github.com/mcorrigan89/media/internal/repositories"
	"github.com/mcorrigan89/media/internal/serviceapis"
	"github.com/mcorrigan89/media/internal/services"

	"github.com/rs/zerolog"
)

type application struct {
	config            config.Config
	wg                *sync.WaitGroup
	logger            *zerolog.Logger
	services          *services.Services
	protoServer       *api.ProtoServer
	serviceApiClients *serviceapis.ServiceApiClients
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

	serviceApiClients := serviceapis.NewServiceApiClients(&cfg, &logger, &wg)
	repositories := repositories.NewRepositories(db, &cfg, &logger, &wg)
	services := services.NewServices(&repositories, &cfg, &logger, &wg)
	protoServer := api.NewProtoServer(&cfg, &logger, &wg, &services)

	app := &application{
		wg:                &wg,
		config:            cfg,
		logger:            &logger,
		services:          &services,
		protoServer:       protoServer,
		serviceApiClients: serviceApiClients,
	}

	err = app.serve()
	if err != nil {
		logger.Err(err).Msg("Failed to start server")
		os.Exit(1)
	}
}
