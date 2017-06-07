package logger

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/Tympanix/artoodetoo/types"
)

// Log is a slice of entries
type Log []*Entry

func (l Log) Len() int {
	return len(l)
}

func (l Log) Less(i, j int) bool {
	return l[i].Time > l[j].Time
}

func (l Log) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

var logs Log
var lock = new(sync.RWMutex)

// Entry is a log entry in the log
type Entry struct {
	Type    string `json:"type"`
	Task    string `json:"task"`
	Message string `json:"message"`
	Time    int64  `json:"time"`
}

// Get retrieved log newer than a specified time
func Get(time int64) Log {
	lock.RLock()
	defer lock.RUnlock()
	log := make(Log, 0)
	for _, l := range logs {
		if l.Time > time {
			log = append(log, l)
		} else {
			break
		}
	}
	_ = sort.Reverse(log)
	return log
}

// Error logs an error to the application
func Error(task types.Identifiable, err error) {
	lock.Lock()
	defer lock.Unlock()
	logs = append(logs, &Entry{
		Type:    "error",
		Task:    task.ID(),
		Message: fmt.Sprint(err),
		Time:    time.Now().Unix(),
	})
}

// Success appends a success message to the log
func Success(task types.Identifiable, message string) {
	lock.Lock()
	defer lock.Unlock()
	logs = append(logs, &Entry{
		Type:    "success",
		Task:    task.ID(),
		Message: message,
		Time:    time.Now().Unix(),
	})
}
