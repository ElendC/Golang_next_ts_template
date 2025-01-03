package authentication

import (
	//"backend/database"
	"encoding/json"
	"fmt"
	//"github.com/jackc/pgx/v5/pgxpool"
	"backend/utils"
	"net/http"
)

type RegisterRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// pool *pgxpool.Pool,
func RegisterHandler(w http.ResponseWriter, req *http.Request) {

	err := utils.Request(w, req, http.MethodPost)
	if err != nil {
		fmt.Println("Error in request validation:", err)
		return
	}

	// Decode and save payload to struct
	var rq RegisterRequest
	err = json.NewDecoder(req.Body).Decode(&rq)
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

	// Preparing response
	contentType := "application/json"
	httpStatus := http.StatusOK
	message := "User registered successfully bg message"
	utils.Response(contentType, httpStatus, message, w)

	fmt.Printf("Registering user %s\n", rq.UserName)
}

// TODO: SAve in db
