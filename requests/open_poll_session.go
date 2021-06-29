package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// OpenPollSession
// https://canvas.instructure.com/doc/api/poll_sessions.html
//
// Path Parameters:
// # Path.PollID (Required) ID
// # Path.ID (Required) ID
//
type OpenPollSession struct {
	Path struct {
		PollID string `json:"poll_id" url:"poll_id,omitempty"` //  (Required)
		ID     string `json:"id" url:"id,omitempty"`           //  (Required)
	} `json:"path"`
}

func (t *OpenPollSession) GetMethod() string {
	return "GET"
}

func (t *OpenPollSession) GetURLPath() string {
	path := "polls/{poll_id}/poll_sessions/{id}/open"
	path = strings.ReplaceAll(path, "{poll_id}", fmt.Sprintf("%v", t.Path.PollID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *OpenPollSession) GetQuery() (string, error) {
	return "", nil
}

func (t *OpenPollSession) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *OpenPollSession) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *OpenPollSession) HasErrors() error {
	errs := []string{}
	if t.Path.PollID == "" {
		errs = append(errs, "'Path.PollID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *OpenPollSession) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
