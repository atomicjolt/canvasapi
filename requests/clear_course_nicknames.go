package requests

import (
	"net/url"

	"github.com/atomicjolt/canvasapi"
)

// ClearCourseNicknames Remove all stored course nicknames.
// https://canvas.instructure.com/doc/api/users.html
//
type ClearCourseNicknames struct {
}

func (t *ClearCourseNicknames) GetMethod() string {
	return "DELETE"
}

func (t *ClearCourseNicknames) GetURLPath() string {
	return ""
}

func (t *ClearCourseNicknames) GetQuery() (string, error) {
	return "", nil
}

func (t *ClearCourseNicknames) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ClearCourseNicknames) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ClearCourseNicknames) HasErrors() error {
	return nil
}

func (t *ClearCourseNicknames) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
