package event_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Tympanix/artoodetoo/assert"
	"github.com/Tympanix/artoodetoo/event"
	"github.com/Tympanix/artoodetoo/plugins/cron"
	"github.com/Tympanix/artoodetoo/task"
	"github.com/Tympanix/artoodetoo/types"
)

type TestEvent struct {
	event.Base
	Name  string `io:"output"`
	start chan struct{}
	stop  chan struct{}
}

func (TestEvent) Describe() string {
	return "An event for testing"
}

func (t *TestEvent) Listen(stop <-chan struct{}) error {
	defer func() {
		t.stop <- struct{}{}
	}()
	t.start <- struct{}{}
	<-stop
	return nil
}

func TestEventSubscribe(t *testing.T) {
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

type dummyRunnable struct {
	fn func(ts types.TupleSpace) error
}

func (d *dummyRunnable) Run(ts types.TupleSpace) error {
	return d.fn(ts)
}

func TestEventTrigger(t *testing.T) {
	ev := new(TestEvent)
	e := event.New(ev)
	called := false

	testRun := func(ts types.TupleSpace) error {
		called = true
		return nil
	}

	e.Subscribe(&dummyRunnable{testRun})
	e.Trigger()
	assert.True(t, called)
}

func TestEventStart(t *testing.T) {
	start := make(chan struct{})
	stop := make(chan struct{})

	ev := &TestEvent{start: start, stop: stop}
	e := event.New(ev)

	e.Start()
	<-start
	e.Start() // NoOp!
	e.Stop()
	<-stop
}

func TestEventProxyUnmarshal(t *testing.T) {
	ev := &TestEvent{}
	e := event.New(ev)

	event.AddEvent(e)

	s := e.ID()
	proxy := event.Proxy{}
	data, _ := json.Marshal(s)
	err := proxy.UnmarshalJSON(data)

	assert.NotError(t, err)
	assert.Equal(t, proxy.Event, e)

	event.RemoveEvent(e)
}

func TestEventProxyMarshal(t *testing.T) {
	ev := &TestEvent{}
	e := event.New(ev)
	proxy := event.Proxy{e}

	data, _ := proxy.MarshalJSON()
	res := fmt.Sprintf(`"%s"`, e.ID())
	assert.Equal(t, string(data), res)
}

func TestEventProxyUnmarshalEvent(t *testing.T) {
	ev := &TestEvent{}
	e := event.New(ev)
	e.AddStatic("Name", "John Doe")

	event.AddEvent(e)

	proxy := event.Proxy{}
	data, _ := json.Marshal(e)
	err := proxy.UnmarshalJSON(data)

	assert.NotError(t, err)
	assert.Equal(t, proxy.Event, e)

	event.RemoveEvent(e)
}
