package requests

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdateCalendarEvent Update and return a calendar event
// https://canvas.instructure.com/doc/api/calendar_events.html
//
// Path Parameters:
// # ID (Required) ID
//
// Form Parameters:
// # CalendarEvent (Optional) Context code of the course/group/user to move this event to.
//    Scheduler appointments and events with section-specific times cannot be moved between calendars.
// # CalendarEvent (Optional) Short title for the calendar event.
// # CalendarEvent (Optional) Longer HTML description of the event.
// # CalendarEvent (Optional) Start date/time of the event.
// # CalendarEvent (Optional) End date/time of the event.
// # CalendarEvent (Optional) Location name of the event.
// # CalendarEvent (Optional) Location address
// # CalendarEvent (Optional) Time zone of the user editing the event. Allowed time zones are
//    {http://www.iana.org/time-zones IANA time zones} or friendlier
//    {http://api.rubyonrails.org/classes/ActiveSupport/TimeZone.html Ruby on Rails time zones}.
// # CalendarEvent (Optional) When true event is considered to span the whole day and times are ignored.
// # CalendarEvent (Optional) Section-level start time(s) if this is a course event. X can be any
//    identifier, provided that it is consistent across the start_at, end_at
//    and context_code
// # CalendarEvent (Optional) Section-level end time(s) if this is a course event.
// # CalendarEvent (Optional) Context code(s) corresponding to the section-level start and end time(s).
//
type UpdateCalendarEvent struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`

	Form struct {
		CalendarEvent struct {
			ContextCode     string                                       `json:"context_code"`     //  (Optional)
			Title           string                                       `json:"title"`            //  (Optional)
			Description     string                                       `json:"description"`      //  (Optional)
			StartAt         time.Time                                    `json:"start_at"`         //  (Optional)
			EndAt           time.Time                                    `json:"end_at"`           //  (Optional)
			LocationName    string                                       `json:"location_name"`    //  (Optional)
			LocationAddress string                                       `json:"location_address"` //  (Optional)
			TimeZoneEdited  string                                       `json:"time_zone_edited"` //  (Optional)
			AllDay          bool                                         `json:"all_day"`          //  (Optional)
			ChildEventData  map[string]UpdateCalendarEventChildEventData `json:"child_event_data"` //  (Optional)
		} `json:"calendar_event"`
	} `json:"form"`
}

func (t *UpdateCalendarEvent) GetMethod() string {
	return "PUT"
}

func (t *UpdateCalendarEvent) GetURLPath() string {
	path := "calendar_events/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateCalendarEvent) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateCalendarEvent) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateCalendarEvent) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateCalendarEvent) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}

type UpdateCalendarEventChildEventData struct {
	StartAt     time.Time `json:"start_at"`     //  (Optional)
	EndAt       time.Time `json:"end_at"`       //  (Optional)
	ContextCode string    `json:"context_code"` //  (Optional)
}
