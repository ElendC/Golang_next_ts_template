package routes

import (
	/*	"backend/handlers"
	 */"backend/handlers/about"
	"backend/handlers/authentication"
	"backend/handlers/home"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/", home.Handler)       // Route for "/"
	http.HandleFunc("/about", about.Handler) // Route for "/about"
	http.HandleFunc("/register", authentication.RegisterHandler)
}
