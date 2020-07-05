package route

import (
	"github.com/gorilla/mux"
	"github.com/milamice62/fakeapi/api/models"
)

func Genres(r *mux.Router) {
	genres := r.PathPrefix("/api/v1/genres").Subrouter()
	genres.HandleFunc("/", models.GetGenres).Methods("GET")
	genres.HandleFunc("/", models.AddGenre).Methods("POST")
	genres.HandleFunc("/{id}", models.UpdateGenre).Methods("PUT")
	genres.HandleFunc("/{id}", models.FindGenre).Methods("GET")
	genres.HandleFunc("/{id}", models.DeleteGenre).Methods("DELETE")
}
