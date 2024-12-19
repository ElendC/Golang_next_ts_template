package main

import (
	"backend/handlers/about"
	"backend/handlers/home"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Handlers
	http.HandleFunc("/home", home.Handler)   // Route for "/home"
	http.HandleFunc("/about", about.Handler) // Route for "/about"

	fmt.Println("Backend Server Starting at 8080...")
	//Starts backend server at port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
