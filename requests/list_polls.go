package requests

import (
	"github.com/atomicjolt/canvasapi"
)

// ListPolls Returns the paginated list of polls for the current user.
// https://canvas.instructure.com/doc/api/polls.html
//
type ListPolls struct {
}

func (t *ListPolls) GetMethod() string {
	return "GET"
}

func (t *ListPolls) GetURLPath() string {
	return ""
}

func (t *ListPolls) GetQuery() (string, error) {
	return "", nil
}

func (t *ListPolls) GetBody() (string, error) {
	return "", nil
}

func (t *ListPolls) HasErrors() error {
	return nil
}

func (t *ListPolls) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
