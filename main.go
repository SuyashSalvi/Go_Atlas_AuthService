package main

import (
	"log"
	"net/http"
	"authservice2/db"
	"authservice2/routes"
	"authservice2/middleware"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database
	db.Init()

	// Create a new router
	router := mux.NewRouter()

	// Apply the middleware
	router.Use(middleware.LoggingMiddleware)
	
	// Initialize routes
	routes.InitRoutes(router)

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
