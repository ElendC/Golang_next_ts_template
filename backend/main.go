package main

import (
	"backend/database"
	"backend/environment"
	"backend/routes"
	"log"
)

func main() {
	err := environment.LoadEnvFile("./environment/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	} //TODO: make a print instead of log

	// Initialize and test Database Connection
	database.DbConfig()

	// Handlers
	routes.SetupRoutes()

	//Starts backend server
	ServerSetup()

}
