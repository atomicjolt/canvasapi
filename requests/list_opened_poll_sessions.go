package requests

import (
	"github.com/atomicjolt/canvasapi"
)

// ListOpenedPollSessions A paginated list of all opened poll sessions available to the current user.
// https://canvas.instructure.com/doc/api/poll_sessions.html
//
type ListOpenedPollSessions struct {
}

func (t *ListOpenedPollSessions) GetMethod() string {
	return "GET"
}

func (t *ListOpenedPollSessions) GetURLPath() string {
	return ""
}

func (t *ListOpenedPollSessions) GetQuery() (string, error) {
	return "", nil
}

func (t *ListOpenedPollSessions) GetBody() (string, error) {
	return "", nil
}

func (t *ListOpenedPollSessions) HasErrors() error {
	return nil
}

func (t *ListOpenedPollSessions) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
