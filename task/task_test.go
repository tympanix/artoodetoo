package task_test

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/Tympanix/automato/assert"
	"github.com/Tympanix/automato/example"
	"github.com/Tympanix/automato/task"
	"github.com/Tympanix/automato/unit"
)

var buf bytes.Buffer

func startCapture() {
	log.SetOutput(&buf)
}

func stopCapture() string {
	log.SetOutput(os.Stderr)
	return buf.String()
}

func defaultTask() (task.Task, *unit.Unit, *unit.Unit) {
	person := unit.NewUnit(&example.PersonEvent{})
	person.SetName("Person")

	email := unit.NewUnit(&example.EmailAction{})
	email.SetName("Email")

	task := task.Task{
		Name:  "myTask",
		Event: person,
		Actions: []*unit.Unit{
			email,
		},
	}
	return task, person, email
}

func TestUnitSameId(t *testing.T) {
	person := unit.NewUnit(&example.PersonEvent{})
	person.SetName("Person")

	subject := unit.NewUnit(&example.StringConverter{})
	subject.SetName("Subject")
	subject.AddStatic("Format", "Good to see you %s")
	subject.AddVar("Placeholder", "Person", "Name")

	message := unit.NewUnit(&example.StringConverter{})
	message.SetName("Message")
	message.AddStatic("Format", "Are your really %v years old?")
	message.AddVar("Placeholder", "Person", "Age")

	email := unit.NewUnit(&example.EmailAction{})
	email.SetName("Email")
	email.AddVar("Subject", "Subject", "Formatted")
	email.AddVar("Message", "Message", "Formatted")
	email.AddStatic("Receiver", "john@doe.com")

	task := task.Task{
		Name:  "myTask",
		Event: person,
		Actions: []*unit.Unit{
			subject,
			message,
			email,
		},
	}

	startCapture()
	err := task.Run()
	stopCapture()

	assert.NotError(t, err)

}

func TestTaskGetUnitByName(t *testing.T) {
	task, person, email := defaultTask()

	p, _ := task.GetUnitByName("Person")
	assert.Equal(t, p, person)

	e, _ := task.GetUnitByName("Email")
	assert.Equal(t, e, email)
}

func TestTaskRegister(t *testing.T) {
	t1, _, _ := defaultTask()
	task.Register(&t1)

	t2, _ := task.GetTaskByName("myTask")

	assert.Equal(t, &t1, t2)

	assert.Equal(t, len(task.All()), 1)
	assert.DeepEqual(t, task.All(), []*task.Task{&t1})

	tnew, _, _ := defaultTask()
	task.Update(&tnew)

	t3, _ := task.GetTaskByName("myTask")
	assert.Equal(t, &tnew, t3)

	task.Unregister(&tnew)
	assert.Equal(t, len(task.All()), 0)
}

func TestTaskUnregisterError(t *testing.T) {
	t1, _, _ := defaultTask()
	err := task.Unregister(&t1)
	assert.ErrorContains(t, err, "unregister", "not found")
}

func TestTaskRegisterError(t *testing.T) {
	t1, _, _ := defaultTask()
	err := task.Register(&t1)
	assert.NotError(t, err)
	err = task.Register(&t1)
	assert.ErrorContains(t, err, "name", "exists")
}

func TestTaskUpdateError(t *testing.T) {
	t1, _, _ := defaultTask()
	t1.Name = "DoesNotExist"
	err := task.Update(&t1)
	assert.ErrorContains(t, err, "update", "not found")
}

func TestTaskDescribe(t *testing.T) {
	task, _, _ := defaultTask()

	startCapture()
	task.Describe()
	output := stopCapture()

	assert.True(t, strings.Contains(output, "myTask"))
	assert.True(t, strings.Contains(output, "example.PersonEvent"))
	assert.True(t, strings.Contains(output, "example.EmailAction"))
}
