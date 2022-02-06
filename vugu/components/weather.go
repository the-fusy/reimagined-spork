package components

import (
	"github.com/spadiff/reimagined-spork/internal/weather"
	"github.com/vugu/vugu"
	"strconv"
	"time"
)

type Weather struct {
	City string

	date    time.Time
	minDate time.Time
	maxDate time.Time

	dayData map[time.Time]weather.DayData
}

func (c *Weather) Init(ctx vugu.InitCtx) {
	c.City = "London"
	c.date = time.Now().UTC().Truncate(24 * time.Hour)
	c.minDate = c.date
	c.maxDate = c.date.Add(2 * 24 * time.Hour)

	go func() {
		dayData, err := weather.GetForecast(c.City)

		ctx.EventEnv().Lock()
		if err != nil {
			c.dayData = nil
		} else {
			c.dayData = dayData
		}
		ctx.EventEnv().UnlockRender()
	}()
}

func (c *Weather) Date() string {
	return c.date.Format("02 January")
}

func (c *Weather) Degree() string {
	if c.dayData == nil {
		return "ERR"
	}
	return strconv.FormatFloat(float64(c.dayData[c.date].Temp), 'f', -1, 32)
}

func (c *Weather) LoadPrevDay(e vugu.DOMEvent) {
	c.date = c.date.Add(-24 * time.Hour)
	if c.date.Before(c.minDate) {
		c.date = c.minDate
	}
}

func (c *Weather) LoadNextDay(e vugu.DOMEvent) {
	c.date = c.date.Add(24 * time.Hour)
	if c.date.After(c.maxDate) {
		c.date = c.maxDate
	}
}
