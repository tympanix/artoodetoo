package task_test

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/Tympanix/artoodetoo/assert"
	"github.com/Tympanix/artoodetoo/event"
	"github.com/Tympanix/artoodetoo/state"
	"github.com/Tympanix/artoodetoo/task"
	"github.com/Tympanix/artoodetoo/unit"
)

type DummyUnit struct {
	StringA string `io:"input"`
	StringB string `io:"input"`

	Result string `io:"output"`
}

// Describe describes what a Concatenation does
func (s *DummyUnit) Describe() string {
	return "Concatenates string A with string B"
}

// Execute function concatenates strings
func (s *DummyUnit) Execute() error {
	s.Result = s.StringA + s.StringB
	return nil
}

type DummyEvent struct {
	event.Base
	Value string `io:"input"`
}

func (d *DummyEvent) Listen(<-chan struct{}) error {
	return nil
}

func (DummyEvent) Describe() string {
	return "A dummy event"
}

var buf bytes.Buffer

func startCapture() {
	log.SetOutput(&buf)
}

func stopCapture() string {
	log.SetOutput(os.Stderr)
	return buf.String()
}

func TestTaskSubscribe(t *testing.T) {
	ac := unit.NewUnit(new(DummyUnit))
	ta := task.New(ac)

	ev := event.New(new(DummyEvent))
	ta.SetEvent(ev)

	ta.Subscribe()
	assert.Equal(t, len(ev.Observers), 1)
	assert.Equal(t, ev.Observers[0], ta)

	ta.Unsubscribe()
	assert.Equal(t, len(ev.Observers), 0)
}

func TestTaskValidateActionNoName(t *testing.T) {
	ac := unit.NewUnit(new(DummyUnit))
	ta := task.New(ac)

	err := ta.Validate()
	assert.Error(t, err)
}

func TestTaskValidate(t *testing.T) {
	ac := unit.NewUnit(new(DummyUnit))
	ta := task.New(ac)

	ac.SetName("MyAction")
	ac.AddStatic("StringA", "John")
	ac.AddStatic("StringB", "Doe")

	err := ta.Validate()
	assert.Error(t, err)
}

func TestTaskValidateHasCycle(t *testing.T) {
	ac := unit.NewUnit(new(DummyUnit))
	ac2 := unit.NewUnit(new(DummyUnit))
	ta := task.New(ac, ac2)

	ev := event.New(new(DummyEvent))
	ta.SetEvent(ev)

	ac.SetName("Action1")
	ac.AddVar("StringA", "Action2", "Result")
	ac.AddStatic("StringB", "John")

	ac2.SetName("Action2")
	ac2.AddVar("StringA", "Action1", "Result")
	ac2.AddStatic("StringB", "Doe")

	err := ta.Validate()
	assert.Error(t, err)
}

func TestTaskGenerateUUID(t *testing.T) {
	ac := unit.NewUnit(new(DummyUnit))
	ta := task.New(ac)

	old := ta.ID()
	ta.GenerateUUID()
	assert.NotEqual(t, old, ta.ID())
}

func TestTaskGenUnitByName(t *testing.T) {
	ac := unit.NewUnit(new(DummyUnit))
	ac2 := unit.NewUnit(new(DummyUnit))
	ta := task.New(ac, ac2)

	ev := event.New(new(DummyEvent))
	ta.SetEvent(ev)

	ac.SetName("Action1")
	ac2.SetName("Action2")

	u, _ := ta.GetUnitByName("Action1")
	assert.Equal(t, u, ac)
	_, err := ta.GetUnitByName("DoesntExist")
	assert.Error(t, err)
}

type relayAction struct {
	Input  interface{} `io:"input"`
	Output interface{} `io:"output"`
	fn     func(*relayAction) error
}

func (relayAction) Describe() string {
	return "Relays an actions to an external function"
}

func (r *relayAction) Execute() error {
	return r.fn(r)
}

func Relay(fn func(*relayAction) error) *unit.Unit {
	return unit.NewUnit(&relayAction{fn: fn})
}

func TestTaskRun(t *testing.T) {
	out := make(chan string)

	ac := Relay(func(a *relayAction) error {
		s := a.Input.(string)
		assert.Equal(t, s, "John Doe")
		a.Output = strings.ToLower(s)
		out <- s
		return nil
	})

	ac2 := Relay(func(a *relayAction) error {
		s := a.Input.(string)
		assert.Equal(t, s, "john doe")
		out <- s
		return nil
	})

	ta := task.New(ac, ac2)

	ev := event.New(new(DummyEvent))
	ta.SetEvent(ev)

	ac.SetName("Action1")
	ac.AddStatic("Input", "John Doe")

	ac2.SetName("Action2")
	ac2.AddVar("Input", "Action1", "Output")

	err := ta.Run(state.New())
	assert.NotError(t, err)

	var s string
	s = <-out
	assert.Equal(t, s, "John Doe")
	s = <-out
	assert.Equal(t, s, "john doe")
}

func TestTaskRegister(t *testing.T) {
	ac := unit.NewUnit(new(DummyUnit))
	ac2 := unit.NewUnit(new(DummyUnit))
	ta := task.New(ac, ac2)
	ta.GenerateUUID()

	ev := event.New(new(DummyEvent))
	ta.SetEvent(ev)
	ta.Name = "MyTask"
	task.AddTask(ta)

	// Add task and get by name and ID
	_ta, err := task.GetTaskByID(ta.ID())
	assert.NotError(t, err)
	assert.Equal(t, _ta, ta)
	_ta, err = task.GetTaskByName(ta.Name)
	assert.NotError(t, err)
	assert.Equal(t, _ta, ta)
	tasks := task.All()
	assert.Equal(t, len(tasks), 1)

	// Copy task, modify and update
	copy := new(task.Task)
	*copy = *ta
	copy.Name = "MyTask2"
	err = task.Update(copy)
	assert.NotError(t, err)
	_ta, err = task.GetTaskByName("MyTask2")
	assert.NotError(t, err)
	assert.Equal(t, _ta, copy)

	// Remove task
	err = task.RemoveTask(ta)
	assert.NotError(t, err)
	tasks = task.All()
	assert.Equal(t, len(tasks), 0)
	err = task.RemoveTask(ta)
	assert.Error(t, err)
}
