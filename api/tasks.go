package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Tympanix/automato/state"
	"github.com/Tympanix/automato/task"
	"github.com/Tympanix/automato/util"
	"github.com/gorilla/mux"
)

func init() {
	api.HandleFunc("/tasks", listTasks).Methods("GET")
	api.HandleFunc("/tasks", newTask).Methods("POST")
	api.HandleFunc("/tasks", updateTask).Methods("PUT")
	api.HandleFunc("/tasks/{task}", deleteTask).Methods("DELETE")
	api.HandleFunc("/tasks/{task}", runTask).Methods("POST")
}

func newTask(w http.ResponseWriter, r *http.Request) {
	var task task.Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	task.GenerateUUID()
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
	taskid, ok := vars["task"]
	if !ok {
		http.Error(w, "No task given", http.StatusInternalServerError)
		return
	}
	t, err := task.GetTaskByID(taskid)
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
	taskid, ok := vars["task"]
	if !ok {
		http.Error(w, "No task given", http.StatusInternalServerError)
		return
	}
	task, err := task.GetTaskByID(taskid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := task.Run(state.New()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
