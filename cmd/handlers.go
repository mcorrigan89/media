package main

import (
	"net/http"
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
	app.logger.Info().Ctx(ctx).Msg("/v1/upload")

	r.ParseMultipartForm(200 << 20)

	file, handler, err := r.FormFile("image")
	if err != nil {
		app.logger.Err(err).Msg("Failed to get file from form")
		http.Error(w, "Failed to get file from form", http.StatusBadRequest)
		return
	}

	defer file.Close()

	err = app.services.StorageService.UploadObject(ctx, handler.Filename, file, handler.Size)
	if err != nil {
		app.logger.Err(err).Msg("Failed to upload object to storage")
		http.Error(w, "Failed to upload object to storage", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Successfully uploaded"))
}
