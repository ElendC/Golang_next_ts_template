package main

import (
	"backend/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Handlers
	routes.SetupRoutes()

	fmt.Println("Backend Server Starting at 8080...")
	//Starts backend server at port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
