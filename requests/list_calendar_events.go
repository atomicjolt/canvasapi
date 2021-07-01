package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// ListCalendarEvents Retrieve the paginated list of calendar events or assignments for the current user
// https://canvas.instructure.com/doc/api/calendar_events.html
//
// Query Parameters:
// # Query.Type (Optional) . Must be one of event, assignmentDefaults to "event"
// # Query.StartDate (Optional) Only return events since the start_date (inclusive).
//    Defaults to today. The value should be formatted as: yyyy-mm-dd or ISO 8601 YYYY-MM-DDTHH:MM:SSZ.
// # Query.EndDate (Optional) Only return events before the end_date (inclusive).
//    Defaults to start_date. The value should be formatted as: yyyy-mm-dd or ISO 8601 YYYY-MM-DDTHH:MM:SSZ.
//    If end_date is the same as start_date, then only events on that day are
//    returned.
// # Query.Undated (Optional) Defaults to false (dated events only).
//    If true, only return undated events and ignore start_date and end_date.
// # Query.AllEvents (Optional) Defaults to false (uses start_date, end_date, and undated criteria).
//    If true, all events are returned, ignoring start_date, end_date, and undated criteria.
// # Query.ContextCodes (Optional) List of context codes of courses/groups/users whose events you want to see.
//    If not specified, defaults to the current user (i.e personal calendar,
//    no course/group events). Limited to 10 context codes, additional ones are
//    ignored. The format of this field is the context type, followed by an
//    underscore, followed by the context id. For example: course_42
// # Query.Excludes (Optional) Array of attributes to exclude. Possible values are "description", "child_events" and "assignment"
//
type ListCalendarEvents struct {
	Query struct {
		Type         string    `json:"type" url:"type,omitempty"`                   //  (Optional) . Must be one of event, assignment
		StartDate    time.Time `json:"start_date" url:"start_date,omitempty"`       //  (Optional)
		EndDate      time.Time `json:"end_date" url:"end_date,omitempty"`           //  (Optional)
		Undated      bool      `json:"undated" url:"undated,omitempty"`             //  (Optional)
		AllEvents    bool      `json:"all_events" url:"all_events,omitempty"`       //  (Optional)
		ContextCodes []string  `json:"context_codes" url:"context_codes,omitempty"` //  (Optional)
		Excludes     []string  `json:"excludes" url:"excludes,omitempty"`           //  (Optional)
	} `json:"query"`
}

func (t *ListCalendarEvents) GetMethod() string {
	return "GET"
}

func (t *ListCalendarEvents) GetURLPath() string {
	return ""
}

func (t *ListCalendarEvents) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListCalendarEvents) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListCalendarEvents) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListCalendarEvents) HasErrors() error {
	errs := []string{}
	if t.Query.Type != "" && !string_utils.Include([]string{"event", "assignment"}, t.Query.Type) {
		errs = append(errs, "Type must be one of event, assignment")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListCalendarEvents) Do(c *canvasapi.Canvas) ([]*models.CalendarEvent, *canvasapi.PagedResource, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.CalendarEvent{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, nil, err
	}

	pagedResource, err := canvasapi.ExtractPagedResource(response.Header)
	if err != nil {
		return nil, nil, err
	}

	return ret, pagedResource, nil
}
