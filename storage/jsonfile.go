package storage

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"

	"github.com/Tympanix/automato/task"
)

const fileMode = 0666

// JSONFile is a store implementation that saves tasks consistently to a json file
type JSONFile struct {
	path  string
	Tasks []*task.Task `json:"tasks"`
}

// NewJSONFile creates a new json file storage type
func NewJSONFile(filepath string) (json *JSONFile, err error) {
	json = &JSONFile{
		path: filepath,
	}
	if json.missing() {
		log.Println("Creating new json file")
		var file *os.File
		file, err = json.create()
		if err != nil {
			return
		}
		defer file.Close()
	}
	return
}

func (j *JSONFile) create() (*os.File, error) {
	return os.Create(j.path)
}

func (j *JSONFile) missing() bool {
	_, err := os.Stat(j.path)
	return os.IsNotExist(err)
}

func (j *JSONFile) open() (*os.File, error) {
	return os.OpenFile(j.path, os.O_RDWR, fileMode)
}

// SaveTask saves the task to the json file
func (j *JSONFile) SaveTask(task *task.Task) error {
	if j.taskExists(task) {
		return errors.New("Task could not be saved, already exists in json file")
	}
	j.appendTask(task)
	if err := j.write(); err != nil {
		return err
	}
	return nil
}

func (j *JSONFile) write() error {
	file, err := j.create()
	defer file.Close()
	if err != nil {
		return err
	}
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
	if err := decoder.Decode(j); err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}
	return j.Tasks, nil
}

// DeleteTask deletes a task from the json file
func (j *JSONFile) DeleteTask(t *task.Task) error {
	if err := j.deleteTask(t); err != nil {
		return err
	}
	if err := j.write(); err != nil {
		return err
	}
	return nil
}

// UpdateTask updates as task and saves it to the json file
func (j *JSONFile) UpdateTask(t *task.Task) error {
	i, err := j.indexOfTask(t)
	if err != nil {
		return err
	}
	j.Tasks[i] = t
	if err := j.write(); err != nil {
		return err
	}
	return nil
}

func (j *JSONFile) indexOfTask(task *task.Task) (int, error) {
	for i, t := range j.Tasks {
		if t.Name == task.Name {
			return i, nil
		}
	}
	return -1, errors.New("Task not found in json storage")
}

func (j *JSONFile) deleteTask(task *task.Task) error {
	i, err := j.indexOfTask(task)
	if err != nil {
		return err
	}
	j.Tasks = append(j.Tasks[:i], j.Tasks[i+1:]...)
	return nil
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
