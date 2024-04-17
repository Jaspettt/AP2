package main

import (
	"fmt"
	"net/http"
	"strconv"
	"task1/internal/data"
	"time"
)

func (app *application) createVinylHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}
	err := app.readJSON(r, &input)
	if err != nil {
		app.errorResponse(w, r, 400, err.Error())
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}
func (app *application) showVinyl(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.errorResponse(w, r, 400, err.Error())
		return
	}
	vinyl := data.Vinyls{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "DAMN",
		Artist:    "Kendrick Lamar",
		Year:      2017,
		Runtime:   96,
		Genres:    []string{"hip-hop"},
	}
	err = app.writeJSON(w, 200, wrap{"vinyl": vinyl})
	if err != nil {
		app.errorResponse(w, r, 500, err.Error())
		return
	}
}
