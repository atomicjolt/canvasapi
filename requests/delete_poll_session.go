package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// DeletePollSession <b>204 No Content</b> response code is returned if the deletion was successful.
// https://canvas.instructure.com/doc/api/poll_sessions.html
//
// Path Parameters:
// # PollID (Required) ID
// # ID (Required) ID
//
type DeletePollSession struct {
	Path struct {
		PollID string `json:"poll_id"` //  (Required)
		ID     string `json:"id"`      //  (Required)
	} `json:"path"`
}

func (t *DeletePollSession) GetMethod() string {
	return "DELETE"
}

func (t *DeletePollSession) GetURLPath() string {
	path := "polls/{poll_id}/poll_sessions/{id}"
	path = strings.ReplaceAll(path, "{poll_id}", fmt.Sprintf("%v", t.Path.PollID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeletePollSession) GetQuery() (string, error) {
	return "", nil
}

func (t *DeletePollSession) GetBody() (string, error) {
	return "", nil
}

func (t *DeletePollSession) HasErrors() error {
	errs := []string{}
	if t.Path.PollID == "" {
		errs = append(errs, "'PollID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeletePollSession) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}