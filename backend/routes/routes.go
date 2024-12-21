package routes

import (
	"backend/handlers/about"
	"backend/handlers/home"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/", home.Handler)       // Route for "/home"
	http.HandleFunc("/about", about.Handler) // Route for "/about"
}
