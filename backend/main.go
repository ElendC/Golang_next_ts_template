package main

import (
	"backend/database"
	"backend/environment"
	"backend/routes"
	"fmt"
)

func main() {
	// Loads environment variables
	err := environment.LoadEnvFile("./environment/.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}

	// Initialize database and create tables
	pool, err2 := database.DbConfig()
	if err2 != nil {
		fmt.Printf("Database config failed: %v\n", err2)
		return
	}
	defer pool.Close()

	// Sets up routes for all handlers
	routes.SetupRoutes()

	//Starts backend server
	ServerSetup()

}
