package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Tympanix/artoodetoo/state"
	"github.com/Tympanix/artoodetoo/task"
	"github.com/Tympanix/artoodetoo/util"
	"github.com/gorilla/mux"
)

func init() {
	API.Handle("/tasks", auth(listTasks)).Methods("GET")
	API.Handle("/tasks", auth(newTask)).Methods("POST")
	API.Handle("/tasks", auth(updateTask)).Methods("PUT")
	API.Handle("/tasks/{task}", auth(deleteTask)).Methods("DELETE")
	API.Handle("/tasks/{task}", auth(runTask)).Methods("POST")
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
