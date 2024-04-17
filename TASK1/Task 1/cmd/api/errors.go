package main

import "net/http"

func (app *application) logError(err error) {
	app.logger.Print(err)
}
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	data := wrap{"message": message}
	err := app.writeJSON(w, 200, data)
	if err != nil {
		app.logError(err)
		w.WriteHeader(500)
	}
}
