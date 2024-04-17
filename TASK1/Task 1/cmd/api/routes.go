package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /vinyls/{id}/", app.showVinyl)
	mux.HandleFunc("POST /vinyls/", app.createVinylHandler)
	return mux
}
