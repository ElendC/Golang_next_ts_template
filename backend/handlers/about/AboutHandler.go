package about

import (
	"backend/models"
	"encoding/json"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	data := models.Page{
		Title: "This is the about handler!",
		Body:  "This body is for about handler!",
	}
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins for testing

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data) // Serialize and send JSON
}
