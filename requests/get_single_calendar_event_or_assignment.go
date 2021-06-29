package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetSingleCalendarEventOrAssignment
// https://canvas.instructure.com/doc/api/calendar_events.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
type GetSingleCalendarEventOrAssignment struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetSingleCalendarEventOrAssignment) GetMethod() string {
	return "GET"
}

func (t *GetSingleCalendarEventOrAssignment) GetURLPath() string {
	path := "calendar_events/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetSingleCalendarEventOrAssignment) GetQuery() (string, error) {
	return "", nil
}

func (t *GetSingleCalendarEventOrAssignment) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSingleCalendarEventOrAssignment) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSingleCalendarEventOrAssignment) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleCalendarEventOrAssignment) Do(c *canvasapi.Canvas) (*models.CalendarEvent, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.CalendarEvent{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
