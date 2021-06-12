package models

type CalendarLink struct {
	Ics string `json:"ics"` // The URL of the calendar in ICS format.Example: https://canvas.instructure.com/feeds/calendars/course_abcdef.ics
}

func (t *CalendarLink) HasError() error {
	return nil
}
