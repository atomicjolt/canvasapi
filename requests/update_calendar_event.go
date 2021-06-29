package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdateCalendarEvent Update and return a calendar event
// https://canvas.instructure.com/doc/api/calendar_events.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.CalendarEvent.ContextCode (Optional) Context code of the course/group/user to move this event to.
//    Scheduler appointments and events with section-specific times cannot be moved between calendars.
// # Form.CalendarEvent.Title (Optional) Short title for the calendar event.
// # Form.CalendarEvent.Description (Optional) Longer HTML description of the event.
// # Form.CalendarEvent.StartAt (Optional) Start date/time of the event.
// # Form.CalendarEvent.EndAt (Optional) End date/time of the event.
// # Form.CalendarEvent.LocationName (Optional) Location name of the event.
// # Form.CalendarEvent.LocationAddress (Optional) Location address
// # Form.CalendarEvent.TimeZoneEdited (Optional) Time zone of the user editing the event. Allowed time zones are
//    {http://www.iana.org/time-zones IANA time zones} or friendlier
//    {http://api.rubyonrails.org/classes/ActiveSupport/TimeZone.html Ruby on Rails time zones}.
// # Form.CalendarEvent.AllDay (Optional) When true event is considered to span the whole day and times are ignored.
// # Form.CalendarEvent (Optional) Section-level start time(s) if this is a course event. X can be any
//    identifier, provided that it is consistent across the start_at, end_at
//    and context_code
// # Form.CalendarEvent (Optional) Section-level end time(s) if this is a course event.
// # Form.CalendarEvent (Optional) Context code(s) corresponding to the section-level start and end time(s).
//
type UpdateCalendarEvent struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		CalendarEvent struct {
			ContextCode     string                                       `json:"context_code" url:"context_code,omitempty"`         //  (Optional)
			Title           string                                       `json:"title" url:"title,omitempty"`                       //  (Optional)
			Description     string                                       `json:"description" url:"description,omitempty"`           //  (Optional)
			StartAt         time.Time                                    `json:"start_at" url:"start_at,omitempty"`                 //  (Optional)
			EndAt           time.Time                                    `json:"end_at" url:"end_at,omitempty"`                     //  (Optional)
			LocationName    string                                       `json:"location_name" url:"location_name,omitempty"`       //  (Optional)
			LocationAddress string                                       `json:"location_address" url:"location_address,omitempty"` //  (Optional)
			TimeZoneEdited  string                                       `json:"time_zone_edited" url:"time_zone_edited,omitempty"` //  (Optional)
			AllDay          bool                                         `json:"all_day" url:"all_day,omitempty"`                   //  (Optional)
			ChildEventData  map[string]UpdateCalendarEventChildEventData `json:"child_event_data" url:"child_event_data,omitempty"` //  (Optional)
		} `json:"calendar_event" url:"calendar_event,omitempty"`
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

func (t *UpdateCalendarEvent) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateCalendarEvent) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateCalendarEvent) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
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
	StartAt     time.Time `json:"start_at" url:"start_at,omitempty"`         //  (Optional)
	EndAt       time.Time `json:"end_at" url:"end_at,omitempty"`             //  (Optional)
	ContextCode string    `json:"context_code" url:"context_code,omitempty"` //  (Optional)
}
