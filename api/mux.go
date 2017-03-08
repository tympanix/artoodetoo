package api

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	queryTask = "task"
)

// API is the server mux for handling API calls
var API = mux.NewRouter()

// SetJSON sets the encoding in the http response to json
func SetJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// QueryTask returns the task name from the query of error if the query was not given
func QueryTask(w http.ResponseWriter, r *http.Request) (string, error) {
	values := r.URL.Query()
	taskname := values.Get(queryTask)
	if len(taskname) == 0 {
		http.Error(w, "No task given", http.StatusInternalServerError)
		return "", errors.New("No task given")
	}
	return taskname, nil
}
