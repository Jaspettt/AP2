package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /vinyl/{id}/", limitMiddleware(app.GetVinylByID))
	mux.HandleFunc("GET /vinyls/", limitMiddleware(app.GetAllVinyls))
	mux.HandleFunc("POST /vinyls/", limitMiddleware(app.CreateVinyl))
	mux.HandleFunc("UPDATE /vinyl/{id}/", limitMiddleware(app.UpdateVinyl))
	mux.HandleFunc("DELETE /vinyls/{id}/", limitMiddleware(app.DeleteVinyl))
	return mux
}
