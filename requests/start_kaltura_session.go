package requests

import (
	"net/url"

	"github.com/atomicjolt/canvasapi"
)

// StartKalturaSession Start a new Kaltura session, so that new media can be recorded and uploaded
// to this Canvas instance's Kaltura instance.
// https://canvas.instructure.com/doc/api/services.html
//
type StartKalturaSession struct {
}

func (t *StartKalturaSession) GetMethod() string {
	return "POST"
}

func (t *StartKalturaSession) GetURLPath() string {
	return ""
}

func (t *StartKalturaSession) GetQuery() (string, error) {
	return "", nil
}

func (t *StartKalturaSession) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *StartKalturaSession) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *StartKalturaSession) HasErrors() error {
	return nil
}

func (t *StartKalturaSession) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
