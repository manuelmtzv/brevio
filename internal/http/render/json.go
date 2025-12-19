package render

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func Empty(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}
