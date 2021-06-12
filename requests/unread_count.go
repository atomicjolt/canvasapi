package requests

import (
	"github.com/atomicjolt/canvasapi"
)

// UnreadCount Get the number of unread conversations for the current user
// https://canvas.instructure.com/doc/api/conversations.html
//
type UnreadCount struct {
}

func (t *UnreadCount) GetMethod() string {
	return "GET"
}

func (t *UnreadCount) GetURLPath() string {
	return ""
}

func (t *UnreadCount) GetQuery() (string, error) {
	return "", nil
}

func (t *UnreadCount) GetBody() (string, error) {
	return "", nil
}

func (t *UnreadCount) HasErrors() error {
	return nil
}

func (t *UnreadCount) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
