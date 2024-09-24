package api

import (
	"context"
	"sync"

	"connectrpc.com/connect"

	"github.com/google/uuid"
	mediav1 "github.com/mcorrigan89/media/gen/serviceapis/media/v1"
	"github.com/mcorrigan89/media/internal/config"
	"github.com/mcorrigan89/media/internal/services"

	"github.com/rs/zerolog"
)

type MediaServerV1 struct {
	config   *config.Config
	wg       *sync.WaitGroup
	logger   *zerolog.Logger
	services *services.Services
}

func newMediaProtoUrlServer(cfg *config.Config, logger *zerolog.Logger, wg *sync.WaitGroup, services *services.Services) *MediaServerV1 {
	return &MediaServerV1{
		config:   cfg,
		wg:       wg,
		logger:   logger,
		services: services,
	}
}

func (s *MediaServerV1) GetPhotoById(ctx context.Context, req *connect.Request[mediav1.GetPhotoByIdRequest]) (*connect.Response[mediav1.GetPhotoByIdResponse], error) {
	photoId := req.Msg.PhotoId

	photoUUID, err := uuid.Parse(photoId)
	if err != nil {
		s.logger.Err(err).Ctx(ctx).Msg("Error parsing photo ID")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	photoEntity, err := s.services.PhotoService.GetPhotoByID(ctx, photoUUID)
	if err != nil {
		s.logger.Err(err).Ctx(ctx).Msg("Error getting photo by ID")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&mediav1.GetPhotoByIdResponse{
		Photo: &mediav1.Photo{
			Id:     photoEntity.ID.String(),
			Url:    photoEntity.UrlSlug(),
			Width:  photoEntity.Width,
			Height: photoEntity.Height,
			Size:   photoEntity.Size,
		},
	})
	res.Header().Set("Media-Version", "v1")
	return res, nil
}
