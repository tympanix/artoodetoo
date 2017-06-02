package storage

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"

	"github.com/Tympanix/artoodetoo/event"
	"github.com/Tympanix/artoodetoo/task"
)

const fileMode = 0666

type jsonCache struct {
	Tasks  []*json.RawMessage `json:"tasks"`
	Events []*json.RawMessage `json:"events"`
}

// JSONFile is a store implementation that saves tasks consistently to a json file
type JSONFile struct {
	path   string
	cache  *jsonCache
	Tasks  []*task.Task   `json:"tasks"`
	Events []*event.Event `json:"events"`
}

// NewJSONFile creates a new json file storage type
func NewJSONFile(filepath string) (j *JSONFile, err error) {
	j = &JSONFile{
		path: filepath,
	}

	file, err := j.init()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var cache *jsonCache
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cache); err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}
	j.cache = cache
	return
}

func (j *JSONFile) init() (file *os.File, err error) {
	if j.missing() {
		log.Println("Creating new json file")
		file, err = j.create()
		if err != nil {
			return
		}
	} else {
		file, err = j.open()
		if err != nil {
			return
		}
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

func (j *JSONFile) eventExists(event *event.Event) bool {
	for _, e := range j.Events {
		if e == event || e.ID() == event.ID() {
			return true
		}
	}
	return false
}

func (j *JSONFile) appendEvent(event *event.Event) {
	j.Events = append(j.Events, event)
}

// SaveEvent saves an event in the json file
func (j *JSONFile) SaveEvent(event *event.Event) error {
	if j.eventExists(event) {
		return errors.New("Event could not be saved, already exists in json file")
	}
	j.appendEvent(event)
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

// GetAllTasks returns all tasks loaded from the json file
func (j *JSONFile) GetAllTasks() (tasks []*task.Task, err error) {
	if j.Tasks == nil && j.cache != nil {
		j.Tasks = make([]*task.Task, len(j.cache.Tasks))
		for i, raw := range j.cache.Tasks {
			var task *task.Task
			if err = json.Unmarshal(*raw, &task); err != nil {
				return
			}
			j.Tasks[i] = task
		}
	}
	return j.Tasks, nil
}

// GetAllEvents returns all events loaded from the json file
func (j *JSONFile) GetAllEvents() (events []*event.Event, err error) {
	if j.Events == nil && j.cache != nil {
		j.Events = make([]*event.Event, len(j.cache.Events))
		for i, raw := range j.cache.Events {
			var event *event.Event
			if err = json.Unmarshal(*raw, &event); err != nil {
				return
			}
			j.Events[i] = event
		}
	}
	return j.Events, nil
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
		if t.UUID == task.UUID {
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
		if t.UUID == task.UUID {
			return true
		}
	}
	return false
}
