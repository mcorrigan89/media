package api

import (
	"net/http"
	"sync"

	"connectrpc.com/grpcreflect"
	mediav1connect "github.com/mcorrigan89/media/gen/serviceapis/media/v1/mediav1connect"
	"github.com/mcorrigan89/media/internal/config"
	"github.com/mcorrigan89/media/internal/services"
	"github.com/rs/zerolog"
)

type ProtoServer struct {
	config        *config.Config
	wg            *sync.WaitGroup
	logger        *zerolog.Logger
	services      *services.Services
	mediaV1Server *MediaServerV1
}

func NewProtoServer(cfg *config.Config, logger *zerolog.Logger, wg *sync.WaitGroup, services *services.Services) *ProtoServer {

	mediaV1Server := newMediaProtoUrlServer(cfg, logger, wg, services)

	return &ProtoServer{
		config:        cfg,
		wg:            wg,
		logger:        logger,
		services:      services,
		mediaV1Server: mediaV1Server,
	}
}

func (s *ProtoServer) Handle(r *http.ServeMux) {

	reflector := grpcreflect.NewStaticReflector(
		"serviceapis.media.v1.PhotoService",
	)

	reflectPath, reflectHandler := grpcreflect.NewHandlerV1(reflector)
	r.Handle(reflectPath, reflectHandler)
	reflectPathAlpha, reflectHandlerAlpha := grpcreflect.NewHandlerV1Alpha(reflector)
	r.Handle(reflectPathAlpha, reflectHandlerAlpha)

	mediaV1Path, mediaV1Handle := mediav1connect.NewPhotoServiceHandler(s.mediaV1Server)
	r.Handle(mediaV1Path, mediaV1Handle)
}
