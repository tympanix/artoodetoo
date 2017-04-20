package cron

import (
	"github.com/Tympanix/automato/event"
	"github.com/robfig/cron"
)

// Cron is an event which triggers on specefic time intervals
type Cron struct {
	event.Base
	Cron *cron.Cron `json:"-"`
	Time string     `json:"cron"`
}

func init() {
	event.Register(&Cron{})
}

// Listen starts the cronjob
func (c *Cron) Listen() error {
	c.Cron = cron.New()
	if err := c.Cron.AddFunc(c.Time, c.Trigger); err != nil {
		return err
	}
	c.Cron.Start()
	return nil
}
