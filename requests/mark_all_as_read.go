package requests

import (
	"github.com/atomicjolt/canvasapi"
)

// MarkAllAsRead Mark all conversations as read.
// https://canvas.instructure.com/doc/api/conversations.html
//
type MarkAllAsRead struct {
}

func (t *MarkAllAsRead) GetMethod() string {
	return "POST"
}

func (t *MarkAllAsRead) GetURLPath() string {
	return ""
}

func (t *MarkAllAsRead) GetQuery() (string, error) {
	return "", nil
}

func (t *MarkAllAsRead) GetBody() (string, error) {
	return "", nil
}

func (t *MarkAllAsRead) HasErrors() error {
	return nil
}

func (t *MarkAllAsRead) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
