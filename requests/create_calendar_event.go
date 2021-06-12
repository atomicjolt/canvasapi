package requests

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// CreateCalendarEvent Create and return a new calendar event
// https://canvas.instructure.com/doc/api/calendar_events.html
//
// Form Parameters:
// # CalendarEvent (Required) Context code of the course/group/user whose calendar this event should be
//    added to.
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
// # CalendarEvent (Optional) Number of times to copy/duplicate the event.  Count cannot exceed 200.
// # CalendarEvent (Optional) Defaults to 1 if duplicate `count` is set.  The interval between the duplicated events.
// # CalendarEvent (Optional) . Must be one of daily, weekly, monthlyDefaults to "weekly".  The frequency at which to duplicate the event
// # CalendarEvent (Optional) Defaults to false.  If set to `true`, an increasing counter number will be appended to the event title
//    when the event is duplicated.  (e.g. Event 1, Event 2, Event 3, etc)
//
type CreateCalendarEvent struct {
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
			ChildEventData  map[string]CreateCalendarEventChildEventData `json:"child_event_data"` //  (Optional)
			Duplicate       struct {
				Count          float64 `json:"count"`           //  (Optional)
				Interval       float64 `json:"interval"`        //  (Optional)
				Frequency      string  `json:"frequency"`       //  (Optional) . Must be one of daily, weekly, monthly
				AppendIterator bool    `json:"append_iterator"` //  (Optional)
			} `json:"duplicate"`
		} `json:"calendar_event"`
	} `json:"form"`
}

func (t *CreateCalendarEvent) GetMethod() string {
	return "POST"
}

func (t *CreateCalendarEvent) GetURLPath() string {
	return ""
}

func (t *CreateCalendarEvent) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateCalendarEvent) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateCalendarEvent) HasErrors() error {
	errs := []string{}
	if t.Form.CalendarEvent.ContextCode == "" {
		errs = append(errs, "'CalendarEvent' is required")
	}
	if !string_utils.Include([]string{"daily", "weekly", "monthly"}, t.Form.CalendarEvent.Duplicate.Frequency) {
		errs = append(errs, "CalendarEvent must be one of daily, weekly, monthly")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateCalendarEvent) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}

type CreateCalendarEventChildEventData struct {
	StartAt     time.Time `json:"start_at"`     //  (Optional)
	EndAt       time.Time `json:"end_at"`       //  (Optional)
	ContextCode string    `json:"context_code"` //  (Optional)
}
