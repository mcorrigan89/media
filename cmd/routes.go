package main

import (
	"net/http"
)

func (app *application) ping(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	app.logger.Info().Ctx(ctx).Msg("/ping")
	w.Write([]byte("OK"))
}

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", app.ping)
	mux.HandleFunc("/v1/image/{filename}", app.processImage)
	mux.HandleFunc("/v1/upload-image", app.uploadImage)
	app.protoServer.Handle(mux)

	return app.recoverPanic(app.enabledCORS(app.contextBuilder(mux)))
}
