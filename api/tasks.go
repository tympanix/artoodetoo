package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Tympanix/automato/task"
	"github.com/Tympanix/automato/util"
	"github.com/gorilla/mux"
)

func init() {
	API.HandleFunc("/tasks", listTasks).Methods("GET")
	API.HandleFunc("/tasks", newTask).Methods("POST")
	API.HandleFunc("/tasks", updateTask).Methods("PUT")
	API.HandleFunc("/tasks/{task}", deleteTask).Methods("DELETE")
	API.HandleFunc("/tasks/{task}", runTask).Methods("POST")
}

func newTask(w http.ResponseWriter, r *http.Request) {
	var task task.Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := util.AddTask(&task); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	task.Describe()
}

func listTasks(w http.ResponseWriter, r *http.Request) {
	SetJSON(w)
	encoder := json.NewEncoder(w)
	encoder.Encode(util.AllTasks())
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskname, ok := vars["task"]
	if !ok {
		http.Error(w, "No task given", http.StatusInternalServerError)
		return
	}
	t, err := task.GetTaskByName(taskname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := util.DeleteTask(t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	var task task.Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := util.UpdateTask(&task); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func runTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskname, ok := vars["task"]
	if !ok {
		http.Error(w, "No task given", http.StatusInternalServerError)
		return
	}
	task, err := task.GetTaskByName(taskname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	task.Run()
}
