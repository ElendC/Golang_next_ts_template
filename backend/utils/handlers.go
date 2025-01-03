package utils

import (
	"backend/handlers"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Response sends a standardized HTTP response from the server to the client.
// It takes 4 arguments:
// - contentType: The content type of the response (e.g., "application/json").
// - httpStatusCode: The HTTP status code to send (e.g., http.StatusOK).
// - message: A string message to include in the response body.
// - w: The http.ResponseWriter to send the response to the client.
func Response(contentType string, httpStatusCode int, message string, w http.ResponseWriter) {
	// Set response content type in the header
	w.Header().Set("Content-Type", contentType) // Usually: "application/json"

	// Write HTTP status code
	w.WriteHeader(httpStatusCode) // This is by default ok.

	// Write message as JSON response in the body to the frontend
	resp := map[string]string{"message": message}
	json.NewEncoder(w).Encode(resp)
}

// Request validates an HTTP request by checking CORS, the HTTP method, and the Content-Type header.
// It takes 3 arguments:
// - w: The http.ResponseWriter to send responses to the client.
// - req: The *http.Request object containing the incoming request details.
// - httpMethod: The allowed HTTP method for this handler (e.g., "POST").
// Returns:
// - `nil` if the request is valid and processing can continue.
// - An `error` if validation fails, along with an appropriate HTTP error response.
func Request(w http.ResponseWriter, req *http.Request, httpMethod string) error {
	// Enable CORS for the handler
	handlers.EnableCors(&w)

	// Browser-enforced client-side check to ensure CORS compliance. Respond with 200 OK to proceed.
	// Called CORS preflight request
	if req.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return nil
	}
	// Check if valid method
	if req.Method != httpMethod {
		http.Error(w, fmt.Sprintf("Only %s method is supported for this handler", httpMethod), http.StatusMethodNotAllowed)
		return errors.New(fmt.Sprintf("only %s method is supported", httpMethod))
	}
	// Checks if the header content type from client is json
	if req.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return errors.New("Content-Type must be application/json")
	}
	return nil
}

// TODO: Add error to json.NewEncoder(w)
