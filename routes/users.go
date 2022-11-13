package routes

import (
	"online-cinema/handlers"
	"online-cinema/pkg/mysql"
	"online-cinema/repositories"

	"github.com/gorilla/mux"
)

func UserRoute(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", h.GetUser).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetuserId).Methods("GET")
	// r.HandleFunc("/user/update/{id}", h.UpdateUser).Methods("PATCH")
	// r.HandleFunc("/user/delete/{id}", h.Delete).Methods("DELETE")

}
