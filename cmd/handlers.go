package main

import (
	"encoding/json"
	"net/http"

	"github.com/mcorrigan89/media/internal/services"
)

func (app *application) processImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	app.logger.Info().Ctx(ctx).Msg("/v1/image/{filename}")

	filenameParam := r.PathValue("filename")
	renditionSize := r.URL.Query().Get("rendition")

	imageBytes, err := app.services.StorageService.GetObject(ctx, filenameParam)
	if err != nil {
		app.logger.Err(err).Msg("Failed to get object from storage")
		http.Error(w, "Failed to get object from storage", http.StatusInternalServerError)
		return
	}

	processedImageBytes, contentType, err := app.services.PhotoService.ProcessImage(ctx, imageBytes, renditionSize)
	if err != nil {
		app.logger.Err(err).Msg("Failed to process image")
		http.Error(w, "Failed to process image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", contentType)

	w.Write(processedImageBytes)
}

func (app *application) uploadImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	app.logger.Info().Ctx(ctx).Msg("/v1/upload-image")

	r.ParseMultipartForm(50 << 20)

	file, handler, err := r.FormFile("image")
	if err != nil {
		app.logger.Err(err).Msg("Failed to get file from form")
		http.Error(w, "Failed to get file from form", http.StatusBadRequest)
		return
	}

	defer file.Close()

	photo, err := app.services.PhotoService.CreatePhoto(ctx, services.CreatePhotoArgs{
		Filename: handler.Filename,
		File:     file,
		Size:     handler.Size,
	})
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to create photo")
		http.Error(w, "Failed to create photo", http.StatusInternalServerError)
		return
	}

	type imageJson struct {
		ID   string `json:"id"`
		Slug string `json:"slug"`
	}

	json, err := json.Marshal(imageJson{
		ID:   photo.ID.String(),
		Slug: photo.UrlSlug(),
	})
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to marshal json")
		http.Error(w, "Failed to marshal json", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}
