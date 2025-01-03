package authentication

import (
	"backend/handlers"
	//"backend/database"
	"encoding/json"
	"fmt"
	//"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

type RegisterRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// pool *pgxpool.Pool,
func RegisterHandler(w http.ResponseWriter, req *http.Request) {
	handlers.EnableCors(&w)
	if req.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Check if valid method
	fmt.Println("Register handler called")
	if req.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported for registering", http.StatusMethodNotAllowed)

		return
	}
	// Checks if the header content type from client is json
	if req.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)

		return
	}

	// Save payload to struct
	var rq RegisterRequest
	err := json.NewDecoder(req.Body).Decode(&rq)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)

		return
	}
	defer req.Body.Close()

	// Ensures both username and password are sent from the client
	if rq.UserName == "" || rq.Password == "" {
		http.Error(w, "Both username and password are required", http.StatusBadRequest)
		return
	}

	// Send an ok status to the frontend
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Printf("Registering user %s\n", rq.UserName)
	fmt.Println("finished")
}

// TODO: SAve in db
