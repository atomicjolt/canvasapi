package requests

import (
	"github.com/atomicjolt/canvasapi"
)

// ListClosedPollSessions A paginated list of all closed poll sessions available to the current user.
// https://canvas.instructure.com/doc/api/poll_sessions.html
//
type ListClosedPollSessions struct {
}

func (t *ListClosedPollSessions) GetMethod() string {
	return "GET"
}

func (t *ListClosedPollSessions) GetURLPath() string {
	return ""
}

func (t *ListClosedPollSessions) GetQuery() (string, error) {
	return "", nil
}

func (t *ListClosedPollSessions) GetBody() (string, error) {
	return "", nil
}

func (t *ListClosedPollSessions) HasErrors() error {
	return nil
}

func (t *ListClosedPollSessions) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
