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

	e := event.New(&cron.Cron{
		Time: "@every 1m",
	})

	data, err := json.Marshal(e)
	assert.NotError(t, err)

	out, err := event.UnmarshalJSON(data)
	assert.NotError(t, err)

	cron, ok := out.(*cron.Cron)

	assert.True(t, ok)
	assert.Equal(t, cron.Time, "@every 1m")
	assert.NotEqual(t, event.Templates[cron.Type()], cron)

}
