package requests

import (
	"net/url"

	"github.com/atomicjolt/canvasapi"
)

// HideAllStreamItems Hide all stream items for the user
// https://canvas.instructure.com/doc/api/users.html
//
type HideAllStreamItems struct {
}

func (t *HideAllStreamItems) GetMethod() string {
	return "DELETE"
}

func (t *HideAllStreamItems) GetURLPath() string {
	return ""
}

func (t *HideAllStreamItems) GetQuery() (string, error) {
	return "", nil
}

func (t *HideAllStreamItems) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *HideAllStreamItems) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *HideAllStreamItems) HasErrors() error {
	return nil
}

func (t *HideAllStreamItems) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
