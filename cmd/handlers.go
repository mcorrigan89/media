package main

import (
	"io"
	"math"
	"net/http"

	"github.com/h2non/bimg"
)

var maxSize = 4080

func calculateDimensions(width, height, maxSize int) (int, int) {
	aspectRatio := float64(width) / float64(height)

	var newWidth, newHeight int

	if width > height {
		newWidth = maxSize
		newHeight = int(math.Round(float64(newWidth) / aspectRatio))
	} else {
		newHeight = maxSize
		newWidth = int(math.Round(float64(newHeight) * aspectRatio))
	}

	return newWidth, newHeight
}

func (app *application) processImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	app.logger.Info().Ctx(ctx).Msg("/v1/image/{filename}")

	filenameParam := r.PathValue("filename")

	body, err := app.storage.GetObject(ctx, filenameParam)
	if err != nil {
		app.logger.Err(err).Msg("Failed to get object from storage")
		http.Error(w, "Failed to get object from storage", http.StatusInternalServerError)
		return
	}
	defer body.Close()

	respByte, err := io.ReadAll(body)
	if err != nil {
		app.logger.Err(err).Msg("Failed to read object from storage")
		http.Error(w, "Failed to read object from storage", http.StatusInternalServerError)
		return
	}

	img := bimg.NewImage(respByte)

	metadata, err := img.Metadata()
	if err != nil {
		app.logger.Err(err).Msg("Failed to get metadata")
		http.Error(w, "Failed to get metadata", http.StatusInternalServerError)
		return
	}

	width, height := calculateDimensions(metadata.Size.Width, metadata.Size.Height, maxSize)

	img2, err := img.Resize(width, height)
	if err != nil {
		app.logger.Err(err).Msg("Failed to rotate image")
		http.Error(w, "Failed to rotate image", http.StatusInternalServerError)
		return
	}

	w.Write(img2)
}
