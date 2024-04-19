package main

import "net/http"

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {

	message, err := app.writeJSON(w, status, message)
	if err != nil {
		w.WriteHeader(500)
	}
}
