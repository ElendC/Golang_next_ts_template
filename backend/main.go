package main

import (
	"backend/database"
	"backend/environment"
	"backend/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	err := environment.LoadEnvFile("./environment/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize and test Database Connection
	database.DbConfig()
	
	// Handlers
	routes.SetupRoutes()

	fmt.Println("Backend Server Starting at 8080...")
	//Starts backend server at port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
