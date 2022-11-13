package routes

import (
	"github.com/gorilla/mux"
	"online-cinema/handlers"
	"online-cinema/pkg/mysql"
	"online-cinema/repositories"
)

func CategoriesRoute(r *mux.Router) {
	categoryRepo := repositories.RepositoryCategies(mysql.DB)
	h := handlers.HandlerCategory(categoryRepo)

	r.HandleFunc("/categorys", h.GetCategory).Methods("GET")
	r.HandleFunc("/category/{id}", h.GetCategoryId).Methods("GET")
	r.HandleFunc("/categorys/create", h.CreateCategory).Methods("POST")
	r.HandleFunc("/category/update/{id}", h.UpdateCategory).Methods("PATCH")
	r.HandleFunc("/category/delete/{id}", h.DeleteCategory).Methods("DELETE")

}
