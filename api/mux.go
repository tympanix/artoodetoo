package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Tympanix/automato/task"
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
		json.NewEncoder(w).Encode(task.Components)
	})

	API.HandleFunc("/newtask", func(w http.ResponseWriter, r *http.Request) {
		var task task.Task
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&task)
		if err != nil {
			log.Printf("Error %v", err)
			return
		}
		task.Describe()
	})
}
