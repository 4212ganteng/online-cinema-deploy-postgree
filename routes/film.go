package routes

import (
	"online-cinema/handlers"
	"online-cinema/pkg/midleware"
	"online-cinema/pkg/mysql"
	"online-cinema/repositories"

	"github.com/gorilla/mux"
)

func FilmsRoute(r *mux.Router) {
	categoryRepo := repositories.RepositoryFilm(mysql.DB)
	h := handlers.HandlerFilm(categoryRepo)

	r.HandleFunc("/films", h.GetFilm).Methods("GET")
	r.HandleFunc("/film/{id}", h.GetFilmId).Methods("GET")
	r.HandleFunc("/film/create", midleware.Auth(midleware.UploadFile(h.CreateFilm))).Methods("POST")
	r.HandleFunc("/film/update/{id}", midleware.Auth(midleware.UploadFile(h.UpdateFilm))).Methods("PATCH")
	r.HandleFunc("/film/delete/{id}", h.DeleteFilm).Methods("DELETE")

}
