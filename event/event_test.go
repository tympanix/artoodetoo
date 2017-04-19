package event_test

import (
	"encoding/json"
	"testing"

	"github.com/Tympanix/automato/assert"
	"github.com/Tympanix/automato/event"
	"github.com/Tympanix/automato/service/cron"
	"github.com/Tympanix/automato/task"
)

func TestEvent(t *testing.T) {
	event := event.Base{}
	task := &task.Task{}

	event.Subscribe(task)
	assert.Equal(t, len(event.Observers), 1)

	event.Unsubscribe(task)
	assert.Equal(t, len(event.Observers), 0)

	err := event.Unsubscribe(task)
	assert.Error(t, err)
}

func TestEventMarshal(t *testing.T) {
	e := &cron.Cron{
		Time: "@every 1m",
	}

	data, err := json.Marshal([]event.Event{e})
	var out []event.Event
	json.Unmarshal(data, &out)

	//fmt.Println(string(data))

	assert.NotError(t, err)
	assert.Equal(t, len(out), 1)

	//assert.Equal(t, copy.Time, "@every 1m")
}
