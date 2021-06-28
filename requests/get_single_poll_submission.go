package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetSinglePollSubmission Returns the poll submission with the given id
// https://canvas.instructure.com/doc/api/poll_submissions.html
//
// Path Parameters:
// # PollID (Required) ID
// # PollSessionID (Required) ID
// # ID (Required) ID
//
type GetSinglePollSubmission struct {
	Path struct {
		PollID        string `json:"poll_id" url:"poll_id,omitempty"`                 //  (Required)
		PollSessionID string `json:"poll_session_id" url:"poll_session_id,omitempty"` //  (Required)
		ID            string `json:"id" url:"id,omitempty"`                           //  (Required)
	} `json:"path"`
}

func (t *GetSinglePollSubmission) GetMethod() string {
	return "GET"
}

func (t *GetSinglePollSubmission) GetURLPath() string {
	path := "polls/{poll_id}/poll_sessions/{poll_session_id}/poll_submissions/{id}"
	path = strings.ReplaceAll(path, "{poll_id}", fmt.Sprintf("%v", t.Path.PollID))
	path = strings.ReplaceAll(path, "{poll_session_id}", fmt.Sprintf("%v", t.Path.PollSessionID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetSinglePollSubmission) GetQuery() (string, error) {
	return "", nil
}

func (t *GetSinglePollSubmission) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSinglePollSubmission) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSinglePollSubmission) HasErrors() error {
	errs := []string{}
	if t.Path.PollID == "" {
		errs = append(errs, "'PollID' is required")
	}
	if t.Path.PollSessionID == "" {
		errs = append(errs, "'PollSessionID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSinglePollSubmission) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
