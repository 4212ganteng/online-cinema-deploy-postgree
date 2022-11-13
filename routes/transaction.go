package routes

import (
	"online-cinema/handlers"
	"online-cinema/pkg/midleware"
	"online-cinema/pkg/mysql"
	"online-cinema/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoute(r *mux.Router) {
	transactionRepo := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepo)

	r.HandleFunc("/transactions", h.FindTransaction).Methods("GET")
	r.HandleFunc("/transaction/user/{id}", midleware.Auth(h.FindTransactionUserId)).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.GetTransactionID).Methods("GET")
	r.HandleFunc("/transaction/create", midleware.Auth(h.CreateTransaction)).Methods("POST")
	// midtrans
	r.HandleFunc("/notification", h.Notification).Methods("POST")
}
