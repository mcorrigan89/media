package main

import (
	"net/http"
)

func (app *application) processImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	app.logger.Info().Ctx(ctx).Msg("/v1/image/{filename}")

	filenameParam := r.PathValue("filename")
	renditionSize := r.URL.Query().Get("rendition")

	app.services.StorageService.GetObject(ctx, filenameParam)

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
