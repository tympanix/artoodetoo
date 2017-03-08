package storage

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/Tympanix/automato/task"
)

// JSONFile is a store implementation that saves tasks consistently to a json file
type JSONFile struct {
	path  string
	Tasks []*task.Task `json:"tasks"`
}

// NewJSONFile creates a new json file storage type
func NewJSONFile(filepath string) (*JSONFile, error) {
	json := &JSONFile{
		path: filepath,
	}
	file, err := json.open()
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return json, err
}

func (j *JSONFile) open() (*os.File, error) {
	return os.Create(j.path)
}

// SaveTask saves the task to the json file
func (j *JSONFile) SaveTask(task *task.Task) error {
	if j.taskExists(task) {
		return errors.New("Task could not be saved, already exists in json file")
	}
	file, err := j.open()
	defer file.Close()
	if err != nil {
		return err
	}
	j.appendTask(task)
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(j)
	return nil
}

// GetAllTasks loads all tasks stored in the json file and returns them
func (j *JSONFile) GetAllTasks() ([]*task.Task, error) {
	file, err := j.open()
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	decoder.Decode(j)
	return j.Tasks, nil
}

func (j *JSONFile) appendTask(task *task.Task) {
	j.Tasks = append(j.Tasks, task)
}

func (j *JSONFile) taskExists(task *task.Task) bool {
	for _, t := range j.Tasks {
		if t.Name == task.Name {
			return true
		}
	}
	return false
}
