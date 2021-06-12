package requests

import (
	"github.com/atomicjolt/canvasapi"
)

// ActivityStreamSummary Returns a summary of the current user's global activity stream.
// https://canvas.instructure.com/doc/api/users.html
//
type ActivityStreamSummary struct {
}

func (t *ActivityStreamSummary) GetMethod() string {
	return "GET"
}

func (t *ActivityStreamSummary) GetURLPath() string {
	return ""
}

func (t *ActivityStreamSummary) GetQuery() (string, error) {
	return "", nil
}

func (t *ActivityStreamSummary) GetBody() (string, error) {
	return "", nil
}

func (t *ActivityStreamSummary) HasErrors() error {
	return nil
}

func (t *ActivityStreamSummary) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
