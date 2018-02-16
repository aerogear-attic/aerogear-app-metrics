package web

import (
	"encoding/json"
	"net/http"
)

func withJSON(w http.ResponseWriter, code int, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)

	return json.NewEncoder(w).Encode(payload)
}
