package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// DeleteCalendarEvent Delete an event from the calendar and return the deleted event
// https://canvas.instructure.com/doc/api/calendar_events.html
//
// Path Parameters:
// # ID (Required) ID
//
// Query Parameters:
// # CancelReason (Optional) Reason for deleting/canceling the event.
//
type DeleteCalendarEvent struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`

	Query struct {
		CancelReason string `json:"cancel_reason"` //  (Optional)
	} `json:"query"`
}

func (t *DeleteCalendarEvent) GetMethod() string {
	return "DELETE"
}

func (t *DeleteCalendarEvent) GetURLPath() string {
	path := "calendar_events/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteCalendarEvent) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *DeleteCalendarEvent) GetBody() (string, error) {
	return "", nil
}

func (t *DeleteCalendarEvent) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteCalendarEvent) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
