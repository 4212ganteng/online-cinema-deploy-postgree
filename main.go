package main

import (
	"fmt"
	"log"
	"net/http"
	"online-cinema/databases"
	"online-cinema/pkg/mysql"
	"online-cinema/routes"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println(" No ENV file found")
	}
	//mysql database init
	mysql.DAtabaseInit()
	//migration
	databases.RunMigrate()
	r := mux.NewRouter()

	//	sub route API
	routes.RoutesInit(r.PathPrefix("/api/v1").Subrouter())

	// uploads path prefix
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	//env

	//	cors
	var allowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var allowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"})
	var allowedOrigins = handlers.AllowedOrigins([]string{"*"})

	var port = os.Getenv("PORT")

	fmt.Println("SERVER Running on Port 8000")
	http.ListenAndServe(":"+port, handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(r))

}
