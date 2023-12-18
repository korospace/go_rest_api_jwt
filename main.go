package main

import (
	"go-rest-api-jwt/configs"
	"go-rest-api-jwt/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	configs.ConnectDB()

	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	log.Println("server running at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
