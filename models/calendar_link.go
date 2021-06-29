package models

type CalendarLink struct {
	Ics string `json:"ics" url:"ics,omitempty"` // The URL of the calendar in ICS format.Example: https://canvas.instructure.com/feeds/calendars/course_abcdef.ics
}

func (t *CalendarLink) HasErrors() error {
	return nil
}
