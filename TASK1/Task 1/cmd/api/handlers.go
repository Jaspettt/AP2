package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"task1/internal/data"
)

var NewVinyl data.Vinyl

func (app *application) mainPage(w http.ResponseWriter, r *http.Request) {

	log.WithFields(logrus.Fields{
		"action": "mainPageHandler",
		"method": r.Method,
		"path":   r.URL.Path,
	}).Info("Handling main page request")
	http.ServeFile(w, r, "index.html")
}
func (app *application) GetAllVinyls(w http.ResponseWriter, r *http.Request) {
	vinyls, err := data.GetAllVinyls()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	app.writeJSON(w, 200, vinyls)

}
func (app *application) GetVinylByID(w http.ResponseWriter, r *http.Request) {
	vinylId := r.PathValue("id")
	id, err := strconv.ParseInt(vinylId, 0, 0)
	if err != nil {
		fmt.Sprintf("Incorrect id %v", err)
	}

	vinyl, err := data.GetVinylByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	app.writeJSON(w, 200, vinyl)
}
func (app *application) CreateVinyl(w http.ResponseWriter, r *http.Request) {
	var vinyl data.Vinyl
	err := app.readJSON(r, &vinyl)
	if err != nil {
		fmt.Sprintf("Incorrect json %v", err)
	}
	if err := data.CreateVinyl(&vinyl); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func (app *application) DeleteVinyl(w http.ResponseWriter, r *http.Request) {
	vinylId := r.PathValue("id")
	id, err := strconv.ParseInt(vinylId, 0, 0)
	if err != nil {
		fmt.Sprintf("Incorrect id %v", err)
	}
	if err := data.DeleteVinyl(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (app *application) UpdateVinyl(w http.ResponseWriter, r *http.Request) {
	updateVinyl := data.Vinyl{}
	app.readJSON(r, &updateVinyl)
	vinylId := r.PathValue("id")
	id, err := strconv.ParseInt(vinylId, 0, 0)
	if err != nil {
		fmt.Sprintf("Incorrect id %v", err)
	}
	if err := data.UpdateVinyl(id, updateVinyl); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
