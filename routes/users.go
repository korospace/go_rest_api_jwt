package routes

import (
	"go-rest-api-jwt/controllers"
	"go-rest-api-jwt/middleware"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	router := r.PathPrefix("/users").Subrouter()

	router.Use(middleware.TokenMiddleware)

	router.HandleFunc("/me", controllers.Me).Methods("GET")
}
