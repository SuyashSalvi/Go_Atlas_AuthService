package routes

import (
	"authservice2/handlers"
	"github.com/gorilla/mux"
	"log"
)

// InitRoutes initializes all routes
func InitRoutes(router *mux.Router) {
	log.Println("Initializing routes...")

	router.HandleFunc("/signup", handlers.Signup).Methods("POST")
	log.Println("Route /signup initialized")

	router.HandleFunc("/login", handlers.Login).Methods("POST")
	log.Println("Route /login initialized")

	router.HandleFunc("/token", handlers.AccessToken).Methods("POST")
	log.Println("Route /token initialized")

}
