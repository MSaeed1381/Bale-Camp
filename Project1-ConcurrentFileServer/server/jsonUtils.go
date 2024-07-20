package server

import (
	"encoding/json"
	"net/http"
)

// JSONResponse writes a JSON response to the client.
func JSONResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	marshaled, _ := json.Marshal(payload)
	w.Write(marshaled)
}
