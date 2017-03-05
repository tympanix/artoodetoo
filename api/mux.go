package api

import (
	"encoding/json"
	"net/http"

	"github.com/Tympanix/automato/hub"
)

// API is the server mux for handling API calls
var API = http.NewServeMux()

// SetJSON sets the encoding in the http response to json
func SetJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func init() {
	API.HandleFunc("/test", func(r http.ResponseWriter, w *http.Request) {
		r.Write([]byte("This is test"))
	})

	API.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		SetJSON(w)
		json.NewEncoder(w).Encode(hub.Components)
	})
}
