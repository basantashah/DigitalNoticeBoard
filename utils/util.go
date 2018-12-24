package utils

import (
	"encoding/json"
	"net/http"
)

// //////////////////// //
/* post and get request */
// //////////////////// //

// Message takes status and message as bool and string
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond uses content type as header of response writer
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
