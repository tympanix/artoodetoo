package unit_test

import (
	"encoding/json"
	"sync"
	"testing"

	"github.com/Tympanix/artoodetoo/assert"
	"github.com/Tympanix/artoodetoo/state"
	"github.com/Tympanix/artoodetoo/unit"
)

// SendEmail mimcs sending an email as an action
type DummyUnit struct {
	A int `io:"input"`
	B int `io:"input"`
	R int `io:"output"`
}

// Describe describes what an email action does
func (a *DummyUnit) Describe() string {
	return "Dummy Division"
}

// Execute sends the email
func (a *DummyUnit) Execute() error {
	a.R = a.A / a.B
	return nil
}

func TestUnitExecute(t *testing.T) {
	d := &DummyUnit{}
	u := unit.NewUnit(d)

	wg := new(sync.WaitGroup)
	er := make(chan error)
	ts := state.New()

	u.AddStatic("A", 10)
	u.AddStatic("B", 5)

	wg.Add(1)
	u.RunAsync(wg, ts, er)
	wg.Wait()

	assert.Equal(t, d.R, 2)
}

func TestUnitError(t *testing.T) {
	d := &DummyUnit{}
	u := unit.NewUnit(d)

	wg := new(sync.WaitGroup)
	er := make(chan error)
	ts := state.New()

	u.AddStatic("A", "Opps!")
	u.AddStatic("B", 5)

	wg.Add(1)
	u.RunAsync(wg, ts, er)
	err := <-er
	assert.Error(t, err)
}

func TestUnitInternalError(t *testing.T) {
	d := &DummyUnit{}
	u := unit.NewUnit(d)

	wg := new(sync.WaitGroup)
	er := make(chan error)
	ts := state.New()

	u.AddStatic("A", 10)
	u.AddStatic("B", 0) // Oops!

	wg.Add(1)
	u.RunAsync(wg, ts, er)
	err := <-er
	assert.Error(t, err)
}

func TestUnitNoName(t *testing.T) {
	d := &DummyUnit{}
	u := unit.NewUnit(d)
	err := u.Validate()
	assert.Error(t, err)
}

func TestReturnAction(t *testing.T) {
	d := &DummyUnit{}
	u := unit.NewUnit(d)
	assert.Equal(t, u.Action(), d)
}

func TestUnitNoInput(t *testing.T) {
	d := &DummyUnit{}
	u := unit.NewUnit(d)
	u.Name = "My Unit"
	err := u.Validate()
	assert.Error(t, err)
}

func TestUnitValidated(t *testing.T) {
	d := &DummyUnit{}
	u := unit.NewUnit(d)
	u.Name = "My Unit"
	u.AddStatic("A", 10)
	u.AddStatic("B", 2)
	err := u.Validate()
	assert.NotError(t, err)
}

func TestActionResolver(t *testing.T) {
	r := new(unit.ActionResolver)
	_, err := r.ResolveSubject("Doesn't exists!")
	assert.Error(t, err)
}

func TestUnitJSON(t *testing.T) {
	d := &DummyUnit{}
	u := unit.NewUnit(d)
	unit.Register(d)

	data, err := json.Marshal(u)
	assert.NotError(t, err)

	_u := new(unit.Unit)
	err = json.Unmarshal(data, _u)
	assert.NotError(t, err)
	unit.Unregister(d)
}
