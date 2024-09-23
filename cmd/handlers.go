package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/h2non/bimg"
)

func (app *application) processImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	app.logger.Info().Ctx(ctx).Msg("/v1/image/{filename}")

	filenameParam := r.PathValue("filename")

	body, err := app.storage.GetObject(filenameParam)
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

	meta, err := img.Metadata()
	if err != nil {
		app.logger.Err(err).Msg("Failed to get metadata")
		http.Error(w, "Failed to get metadata", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Image type: %s\n", meta.Type)
	fmt.Printf("Image size: %d\n", meta.Size)
	fmt.Printf("Image width: %d\n", meta.Size.Width)
	fmt.Printf("Image height: %d\n", meta.Size.Height)
	fmt.Printf("Image has EXIF Make: %s\n", meta.EXIF.Make)
	fmt.Printf("Image has EXIF Model: %s\n", meta.EXIF.Model)
	fmt.Printf("Image has EXIF Datetime: %s\n", meta.EXIF.Datetime)

	img2, err := img.Resize(2000, 3000)
	if err != nil {
		app.logger.Err(err).Msg("Failed to rotate image")
		http.Error(w, "Failed to rotate image", http.StatusInternalServerError)
		return
	}

	w.Write(img2)
}
