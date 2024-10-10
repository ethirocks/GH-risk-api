// store common structs
package common

import (
	"encoding/json"
	"net/http"
)

// JSONResponse represents a standard API response format.
type JSONResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// RespondWithJSON sends a JSON response.
func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

// RespondWithError sends an error response.
func RespondWithError(w http.ResponseWriter, status int, message string) {
	RespondWithJSON(w, status, JSONResponse{Success: false, Error: message})
}
