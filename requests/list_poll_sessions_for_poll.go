package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ListPollSessionsForPoll Returns the paginated list of PollSessions in this poll.
// https://canvas.instructure.com/doc/api/poll_sessions.html
//
// Path Parameters:
// # PollID (Required) ID
//
type ListPollSessionsForPoll struct {
	Path struct {
		PollID string `json:"poll_id"` //  (Required)
	} `json:"path"`
}

func (t *ListPollSessionsForPoll) GetMethod() string {
	return "GET"
}

func (t *ListPollSessionsForPoll) GetURLPath() string {
	path := "polls/{poll_id}/poll_sessions"
	path = strings.ReplaceAll(path, "{poll_id}", fmt.Sprintf("%v", t.Path.PollID))
	return path
}

func (t *ListPollSessionsForPoll) GetQuery() (string, error) {
	return "", nil
}

func (t *ListPollSessionsForPoll) GetBody() (string, error) {
	return "", nil
}

func (t *ListPollSessionsForPoll) HasErrors() error {
	errs := []string{}
	if t.Path.PollID == "" {
		errs = append(errs, "'PollID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListPollSessionsForPoll) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
