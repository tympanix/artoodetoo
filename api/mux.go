package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Tympanix/automato/task"
	"github.com/Tympanix/automato/unit"
	"github.com/Tympanix/automato/util"
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
		var metas []*unit.Meta
		for _, v := range unit.Metas {
			metas = append(metas, v)
		}
		json.NewEncoder(w).Encode(metas)
	})

	API.HandleFunc("/newtask", func(w http.ResponseWriter, r *http.Request) {
		var task task.Task
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&task); err != nil {
			log.Printf("Error %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := util.AddTask(&task); err != nil {
			log.Printf("Error %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		task.Describe()
	})

	API.HandleFunc("/runtask", func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		taskname := values.Get("task")
		if len(taskname) == 0 {
			http.Error(w, "No task given", http.StatusInternalServerError)
			return
		}
		task, err := task.GetTaskByName(taskname)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		task.Run()
	})
}
