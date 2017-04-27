package cron

import (
	"errors"

	"github.com/Tympanix/automato/event"
	"github.com/robfig/cron"
)

// Cron is an event which triggers on specefic time intervals
type Cron struct {
	event.Base
	Cron *cron.Cron `json:"-"`
	Time Time       `io:"input"`
}

// Time is a string which describes intervals using the cron spec
type Time string

func init() {
	event.Register(&Cron{})
}

// Describe returns a readable description of the cron job
func (c *Cron) Describe() string {
	return "Cron is a scheduler which triggers on set time intervals"
}

// Listen starts the cronjob
func (c *Cron) Listen() error {
	if len(c.Time) == 0 {
		return errors.New("No time specefied for crontab")
	}
	c.Cron = cron.New()
	if err := c.Cron.AddFunc(string(c.Time), c.Fire); err != nil {
		return err
	}
	c.Cron.Start()
	return nil
}
