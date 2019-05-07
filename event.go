package main

import (
	"github.com/labstack/echo"
	"time"
)

type Event struct {
	Title     string
	Desc      string
	URL       string
	Start     time.Time
	End       time.Time
	TZID      string
	Organizer string
}

func (event *Event) getFormattedStart() string {
	return formatDate(event.Start)
}

func (event *Event) getFormattedEnd() string {
	return formatDate(event.End)
}

func parseDate(input string) (time.Time, error) {
	start, err := time.Parse("2006-01-02T15:04", input)
	return start, err
}

func FromEchoContext(c echo.Context) (*Event, error) {

	start, err := parseDate(c.QueryParam("start"))
	end, err := parseDate(c.QueryParam("end"))

	if err != nil {
		return nil, err
	}

	event := new(Event)

	event.Title = c.QueryParam("title")
	event.URL = c.QueryParam("url")
	event.TZID = c.QueryParam("tzid")
	event.Organizer = c.QueryParam("organizer")
	event.Desc = c.QueryParam("desc")

	event.Start = start
	event.End = end

	return event, nil
}

func formatDate(t time.Time) string {
	return t.Format("20060102T150405")
}
