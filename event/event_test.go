package event_test

import (
	"encoding/json"
	"testing"

	"github.com/Tympanix/artoodetoo/assert"
	"github.com/Tympanix/artoodetoo/event"
	"github.com/Tympanix/artoodetoo/plugins/cron"
	"github.com/Tympanix/artoodetoo/task"
)

func TestEvent(t *testing.T) {
	event := event.New(&cron.Cron{})
	task := &task.Task{}

	event.Subscribe(task)
	assert.Equal(t, len(event.Observers), 1)

	event.Unsubscribe(task)
	assert.Equal(t, len(event.Observers), 0)

	err := event.Unsubscribe(task)
	assert.Error(t, err)
}

func TestEventMarshal(t *testing.T) {

	cron := &cron.Cron{}
	e := event.New(cron)

	e.AddStatic("Time", "@every 1s")

	data, err := json.Marshal(e)
	assert.NotError(t, err)

	var out event.Event
	err = json.Unmarshal(data, &out)
	assert.NotError(t, err)

	assert.NotEqual(t, event.Templates[e.Type()], e)

}
