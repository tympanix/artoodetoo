package logger

import (
	"sort"
	"sync"
	"time"

	"github.com/Tympanix/artoodetoo/types"
)

// History is a slice of entries
type History []*Entry

func (l History) Len() int {
	return len(l)
}

func (l History) Less(i, j int) bool {
	return l[i].Time > l[j].Time
}

func (l History) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

var logs History
var lock = new(sync.RWMutex)

// Entry is a log entry in the log
type Entry struct {
	Type    string `json:"type"`
	Task    string `json:"task"`
	Event   string `json:"event"`
	Message string `json:"message"`
	Time    int64  `json:"time"`
}

func (e *Entry) Error() string {
	return e.Message
}

// SetTask binds a task to the given error
func (e *Entry) SetTask(t types.Identifiable) *Entry {
	e.Task = t.ID()
	return e
}

// SetEvent bind an event to the given error
func (e *Entry) SetEvent(t types.Identifiable) *Entry {
	e.Event = t.ID()
	return e
}

// Log appends the entry to the log history
func (e *Entry) Log() {
	Log(e)
}

// Convert converts an error intro a log entry
func Convert(err error) *Entry {
	return getOrCreate(err)
}

// NewError returns a new entry for logging
func NewError(err string) *Entry {
	return &Entry{
		Type:    "error",
		Message: err,
		Time:    time.Now().Unix(),
	}
}

// NewSuccess returns a new successful log entry
func NewSuccess(err string) *Entry {
	return &Entry{
		Type:    "success",
		Message: err,
		Time:    time.Now().Unix(),
	}
}

func getOrCreate(err error) *Entry {
	if e, ok := err.(*Entry); ok {
		return e
	}
	return NewError(err.Error())
}

// Log logs a new error to the history
func Log(err error) {
	lock.RLock()
	defer lock.RUnlock()
	logs = append(logs, getOrCreate(err))
}

// Get retrieved log newer than a specified time
func Get(time int64) History {
	lock.RLock()
	defer lock.RUnlock()
	log := make(History, 0)
	for i := len(logs) - 1; i >= 0; i-- {
		if logs[i].Time > time {
			log = append(log, logs[i])
		} else {
			break
		}
	}
	_ = sort.Reverse(log)
	return log
}

// Clear clears the log
func Clear() {
	lock.Lock()
	lock.Unlock()
	logs = make(History, 0)
}
