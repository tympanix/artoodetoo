package cron

import (
	"github.com/Tympanix/automato/event"
	"github.com/robfig/cron"
)

type Cron struct {
	event.Base
	Cron *cron.Cron `json:"-"`
	Time string     `json:"cron"`
}

func init() {
	event.Register(&Cron{})
}

func (c *Cron) Listen() error {
	c.Cron = cron.New()
	c.Cron.AddFunc(c.Time, c.Trigger)
	c.Cron.Start()
	return nil
}
