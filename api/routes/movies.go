package route

import (
	"github.com/gorilla/mux"
	"github.com/milamice62/fakeapi/api/models"
)

func Movies(r *mux.Router) {
	movies := r.PathPrefix("/api/v1/movies").Subrouter()
	movies.HandleFunc("/", models.GetMovies).Methods("GET")
	movies.HandleFunc("/", models.AddMovie).Methods("POST")
	movies.HandleFunc("/{id}", models.UpdateMovie).Methods("PUT")
	movies.HandleFunc("/{id}", models.FindMovie).Methods("GET")
	movies.HandleFunc("/{id}", models.DeleteMovie).Methods("DELETE")
}
