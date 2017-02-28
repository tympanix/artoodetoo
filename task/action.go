package task

import "log"

// IAction interface
type IAction interface {
	Execute()
	GetInput() *interface{}
}

// Action interface can operate on input data and perform an action whith that data
type Action struct {
	ID    string      `json:"id"`
	Input interface{} `json:"input"`
}

// Execute executes the action
func (a *Action) Execute() {
	log.Fatal("Execute has not been implemented yet")
}

// GetInput return the input struct
func (a *Action) GetInput() *interface{} {
	return &a.Input
}

// SetID sets a new identifier for the action
func (a *Action) SetID(id string) {
	a.ID = id
}

// GetID gets the current identifier for the action
func (a *Action) GetID() string {
	return a.ID
}
