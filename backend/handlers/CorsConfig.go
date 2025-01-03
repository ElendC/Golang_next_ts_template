package handlers

import (
	"fmt"
	"net/http"
)

// TODO: Do not need in production as frontend most likely will be compiled and serve the backend
// TODO: Can also add this globally, so we won't need to add it in each handler that uses cors
// Tutorial: https://www.stackhawk.com/blog/golang-cors-guide-what-it-is-and-how-to-enable-it/

func EnableCors(w *http.ResponseWriter) {
	fmt.Println("Enabling Cors")
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000/")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}
