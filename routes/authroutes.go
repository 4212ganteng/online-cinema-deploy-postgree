package routes

import (
	"online-cinema/handlers"
	"online-cinema/pkg/midleware"
	"online-cinema/pkg/mysql"
	"online-cinema/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	authRepository := repositories.ReposiitoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
	// for cek auth
	r.HandleFunc("/check-auth", midleware.Auth(h.CheckAuth)).Methods("GET")

}
