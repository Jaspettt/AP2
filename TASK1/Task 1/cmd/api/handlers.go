package main

import "net/http"

func (app *application) example(w http.ResponseWriter, r *http.Request) error {
	data := map[string]string{
		"nice": "cool",
	}
	err := app.writeJSON(w, 200, data)
	if err != nil {
		return err
	}
	return nil
}
