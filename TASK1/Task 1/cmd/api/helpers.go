package main

import (
	"encoding/json"
	"net/http"
)

type wrap map[string]any

func (app *application) writeJSON(w http.ResponseWriter, status int, data any) ([]byte, error) {
	js, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return js, nil
}
func (app *application) readJSON(r *http.Request, data any) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		return err
	}
	return nil
}
