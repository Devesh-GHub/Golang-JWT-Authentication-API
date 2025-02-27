package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	router "github.com/devesh/mongoapi/Routers"
)

func main() {
	fmt.Printf("MongoDB API\n")
	r := router.Router() // Initialize Gorilla Mux router
	log.Println("Server is getting started ...")

	// Set up the server port
	port := os.Getenv("PORT")    //retrieves the value of the PORT environment variable
	if port == "" {              
		port = "8000" // Default port
	}

	router.AuthRoutes(r)
	// Define additional routes
	r.HandleFunc("/api-l", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": "Access granted for api-1"}`))
	}).Methods("GET")

	
	// Start the server
	log.Fatal(http.ListenAndServe(":"+port, r))
}
