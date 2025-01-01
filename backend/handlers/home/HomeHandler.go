package home

import (
	"backend/handlers"
	"backend/models"
	"encoding/json"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	handlers.EnableCors(&w)
	if req.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	data := models.Page{
		Title: "This title is from the backend!",
		Body:  "This body is from the backend!!",
	}
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins for testing

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data) // Serialize and send JSON
}
