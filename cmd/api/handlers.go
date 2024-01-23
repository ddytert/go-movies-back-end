package main

import (
	"backend/cmd/api/internal/models"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	payload := models.Movie{
		ID:          1234,
		Title:       "James Bond - Golden Eye",
		ReleaseDate: time.Now(),
		RunTime:     122,
		MPAARating:  "R",
		Description: "Spy Action",
		Image:       "N/A",
	}
	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, movies)
}
