package cron

import (
	"errors"

	"github.com/Tympanix/automato/event"
	"github.com/robfig/cron"
)

// Cron is an event which triggers on specefic time intervals
type Cron struct {
	event.Base
	Cron  *cron.Cron `json:"-"`
	input struct {
		Time string
	}
}

func init() {
	event.Register(&Cron{})
}

// Describe returns a readable description of the cron job
func (c *Cron) Describe() string {
	return "Cron is a scheduler which triggers on set time intervals"
}

// Input returns the input for the cron job
func (c *Cron) Input() interface{} {
	return &c.input
}

// Output returns the output of the cron job
func (c *Cron) Output() interface{} {
	return nil
}

// Listen starts the cronjob
func (c *Cron) Listen() error {
	if len(c.input.Time) == 0 {
		return errors.New("No time specefied for crontab")
	}
	c.Cron = cron.New()
	if err := c.Cron.AddFunc(c.input.Time, c.Trigger); err != nil {
		return err
	}
	c.Cron.Start()
	return nil
}
