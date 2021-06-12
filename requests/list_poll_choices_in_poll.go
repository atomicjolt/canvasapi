package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ListPollChoicesInPoll Returns the paginated list of PollChoices in this poll.
// https://canvas.instructure.com/doc/api/poll_choices.html
//
// Path Parameters:
// # PollID (Required) ID
//
type ListPollChoicesInPoll struct {
	Path struct {
		PollID string `json:"poll_id"` //  (Required)
	} `json:"path"`
}

func (t *ListPollChoicesInPoll) GetMethod() string {
	return "GET"
}

func (t *ListPollChoicesInPoll) GetURLPath() string {
	path := "polls/{poll_id}/poll_choices"
	path = strings.ReplaceAll(path, "{poll_id}", fmt.Sprintf("%v", t.Path.PollID))
	return path
}

func (t *ListPollChoicesInPoll) GetQuery() (string, error) {
	return "", nil
}

func (t *ListPollChoicesInPoll) GetBody() (string, error) {
	return "", nil
}

func (t *ListPollChoicesInPoll) HasErrors() error {
	errs := []string{}
	if t.Path.PollID == "" {
		errs = append(errs, "'PollID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListPollChoicesInPoll) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
