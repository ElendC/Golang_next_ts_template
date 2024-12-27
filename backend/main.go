package main

import (
	"backend/routes"
)

func main() {

	// Handlers
	routes.SetupRoutes()

	//Starts backend server
	ServerSetup()

}
