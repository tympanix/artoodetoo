package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Tympanix/automato/task"
	"github.com/Tympanix/automato/unit"
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

	API.HandleFunc("/units", func(w http.ResponseWriter, r *http.Request) {
		SetJSON(w)
		var units []*unit.Unit
		for _, v := range unit.Units {
			units = append(units, v)
		}
		json.NewEncoder(w).Encode(units)
	})

	API.HandleFunc("/newtask", func(w http.ResponseWriter, r *http.Request) {
		var task task.Task
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&task)
		if err != nil {
			log.Printf("Error %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		task.Describe()
		task.Run()
	})
}
