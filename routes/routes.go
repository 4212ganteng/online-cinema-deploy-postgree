package routes

import "github.com/gorilla/mux"

func RoutesInit(r *mux.Router) {
	AuthRoutes(r)
	UserRoute(r)
	CategoriesRoute(r)
	FilmsRoute(r)
	TransactionRoute(r)
}
