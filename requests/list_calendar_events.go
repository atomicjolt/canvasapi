package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
// # Type (Optional) . Must be one of event, assignmentDefaults to "event"
// # StartDate (Optional) Only return events since the start_date (inclusive).
//    Defaults to today. The value should be formatted as: yyyy-mm-dd or ISO 8601 YYYY-MM-DDTHH:MM:SSZ.
// # EndDate (Optional) Only return events before the end_date (inclusive).
//    Defaults to start_date. The value should be formatted as: yyyy-mm-dd or ISO 8601 YYYY-MM-DDTHH:MM:SSZ.
//    If end_date is the same as start_date, then only events on that day are
//    returned.
// # Undated (Optional) Defaults to false (dated events only).
//    If true, only return undated events and ignore start_date and end_date.
// # AllEvents (Optional) Defaults to false (uses start_date, end_date, and undated criteria).
//    If true, all events are returned, ignoring start_date, end_date, and undated criteria.
// # ContextCodes (Optional) List of context codes of courses/groups/users whose events you want to see.
//    If not specified, defaults to the current user (i.e personal calendar,
//    no course/group events). Limited to 10 context codes, additional ones are
//    ignored. The format of this field is the context type, followed by an
//    underscore, followed by the context id. For example: course_42
// # Excludes (Optional) Array of attributes to exclude. Possible values are "description", "child_events" and "assignment"
//
type ListCalendarEvents struct {
	Query struct {
		Type         string    `json:"type"`          //  (Optional) . Must be one of event, assignment
		StartDate    time.Time `json:"start_date"`    //  (Optional)
		EndDate      time.Time `json:"end_date"`      //  (Optional)
		Undated      bool      `json:"undated"`       //  (Optional)
		AllEvents    bool      `json:"all_events"`    //  (Optional)
		ContextCodes []string  `json:"context_codes"` //  (Optional)
		Excludes     []string  `json:"excludes"`      //  (Optional)
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
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListCalendarEvents) GetBody() (string, error) {
	return "", nil
}

func (t *ListCalendarEvents) HasErrors() error {
	errs := []string{}
	if !string_utils.Include([]string{"event", "assignment"}, t.Query.Type) {
		errs = append(errs, "Type must be one of event, assignment")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListCalendarEvents) Do(c *canvasapi.Canvas) ([]*models.CalendarEvent, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.CalendarEvent{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}