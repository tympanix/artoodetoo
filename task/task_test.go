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

func TestTaskMissingArgument(t *testing.T) {
	task, _, _ := defaultTask()

	err := task.Validate()
	assert.ErrorContains(t, err, "argument", "missing")
}

func TestTaskUnknownAttribute(t *testing.T) {
	task, _, email := defaultTask()

	email.AddStatic("Receiver", "...")
	email.AddStatic("Message", "...")
	email.AddStatic("Subject", "...")
	email.AddStatic("Lalala", "...") // Oops!

	err := task.Validate()
	assert.ErrorContains(t, err, "unknown", "attribute")
}

func TestTaskUnknownSource(t *testing.T) {
	task, _, email := defaultTask()

	email.AddStatic("Receiver", "...")
	email.AddStatic("Message", "...")
	email.AddVar("Subject", "Missing?!", "...") // Oops!

	err := task.Validate()
	assert.ErrorContains(t, err, "unknown", "source")
}

func TestTaskValueNotString(t *testing.T) {
	task, _, email := defaultTask()

	email.AddStatic("Receiver", "...")
	email.AddStatic("Message", "...")
	email.AddIngredient(unit.Ingredient{
		Type:     unit.IngredientVar,
		Argument: "Subject",
		Source:   "Person",
		Value:    42,
	})

	err := task.Validate()
	assert.ErrorContains(t, err, "invalid", "string")
}

func TestTaskUnknownVariable(t *testing.T) {
	task, _, email := defaultTask()

	email.AddStatic("Receiver", "...")
	email.AddStatic("Message", "...")
	email.AddVar("Subject", "Person", "BlahBlah") // Oops!

	err := task.Validate()
	assert.ErrorContains(t, err, "unknown", "variable")
}

func TestTaskUnknownIngredientType(t *testing.T) {
	task, _, email := defaultTask()

	email.AddStatic("Receiver", "...")
	email.AddStatic("Message", "...")
	email.AddIngredient(unit.Ingredient{
		Type:     42,
		Argument: "Subject",
		Source:   "Person",
		Value:    "Name",
	})

	err := task.Validate()
	assert.ErrorContains(t, err, "Unknown", "ingredient")
}

func TestTaskIngredientUnassignable(t *testing.T) {
	task, _, email := defaultTask()

	email.AddStatic("Receiver", "...")
	email.AddStatic("Message", "...")
	email.AddStatic("Subject", 42) // Oops!

	err := task.Validate()
	assert.ErrorContains(t, err, "type", "incompatible")
}

func TestTaskSuccess(t *testing.T) {
	task, _, email := defaultTask()

	email.AddStatic("Receiver", "my@mail.com")
	email.AddVar("Message", "Person", "Name")
	email.AddStatic("Subject", "Hello World!")

	err := task.Validate()
	assert.NotError(t, err)

	startCapture()
	err = task.Run()
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

func TestTaskRunInputFail(t *testing.T) {
	task, _, email := defaultTask()

	email.AddStatic("Receiver", "my@mail.com")
	email.AddVar("Message", "Person", "Name")
	email.AddStatic("Subject", 42) // Oops!

	startCapture()
	err := task.Run()
	stopCapture()

	assert.ErrorContains(t, err, "Subject", "int")
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
