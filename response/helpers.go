package response

import (
	"encoding/json"
	"net/http"
)

// JSON Response
func JSON(w http.ResponseWriter, i interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(i)
}
