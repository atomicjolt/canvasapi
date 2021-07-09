package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// ListCalendarEventsForUser Retrieve the paginated list of calendar events or assignments for the specified user.
// To view calendar events for a user other than yourself,
// you must either be an observer of that user or an administrator.
// https://canvas.instructure.com/doc/api/calendar_events.html
//
// Path Parameters:
// # Path.UserID (Required) ID
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
// # Query.SubmissionTypes (Optional) When type is "assignment", specifies the allowable submission types for returned assignments.
//    Ignored if type is not "assignment" or if exclude_submission_types is provided.
// # Query.ExcludeSubmissionTypes (Optional) When type is "assignment", specifies the submission types to be excluded from the returned
//    assignments. Ignored if type is not "assignment".
// # Query.ImportantDates (Optional) Defaults to false
//    If true, only events with important dates set to true will be returned. This requires the Important Dates
//    feature flag to be turned on or it will be ignored.
//
type ListCalendarEventsForUser struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Type                   string    `json:"type" url:"type,omitempty"`                                         //  (Optional) . Must be one of event, assignment
		StartDate              time.Time `json:"start_date" url:"start_date,omitempty"`                             //  (Optional)
		EndDate                time.Time `json:"end_date" url:"end_date,omitempty"`                                 //  (Optional)
		Undated                bool      `json:"undated" url:"undated,omitempty"`                                   //  (Optional)
		AllEvents              bool      `json:"all_events" url:"all_events,omitempty"`                             //  (Optional)
		ContextCodes           []string  `json:"context_codes" url:"context_codes,omitempty"`                       //  (Optional)
		Excludes               []string  `json:"excludes" url:"excludes,omitempty"`                                 //  (Optional)
		SubmissionTypes        []string  `json:"submission_types" url:"submission_types,omitempty"`                 //  (Optional)
		ExcludeSubmissionTypes []string  `json:"exclude_submission_types" url:"exclude_submission_types,omitempty"` //  (Optional)
		ImportantDates         bool      `json:"important_dates" url:"important_dates,omitempty"`                   //  (Optional)
	} `json:"query"`
}

func (t *ListCalendarEventsForUser) GetMethod() string {
	return "GET"
}

func (t *ListCalendarEventsForUser) GetURLPath() string {
	path := "users/{user_id}/calendar_events"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *ListCalendarEventsForUser) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListCalendarEventsForUser) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListCalendarEventsForUser) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListCalendarEventsForUser) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if t.Query.Type != "" && !string_utils.Include([]string{"event", "assignment"}, t.Query.Type) {
		errs = append(errs, "Type must be one of event, assignment")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListCalendarEventsForUser) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.CalendarEvent, *canvasapi.PagedResource, error) {
	var err error
	var response *http.Response
	if next != nil {
		response, err = c.Send(next, t.GetMethod(), nil)
	} else {
		response, err = c.SendRequest(t)
	}

	if err != nil {
		return nil, nil, err
	}
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
