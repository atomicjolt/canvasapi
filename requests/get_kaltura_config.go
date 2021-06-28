package requests

import (
	"net/url"

	"github.com/atomicjolt/canvasapi"
)

// GetKalturaConfig Return the config information for the Kaltura plugin in json format.
// https://canvas.instructure.com/doc/api/services.html
//
type GetKalturaConfig struct {
}

func (t *GetKalturaConfig) GetMethod() string {
	return "GET"
}

func (t *GetKalturaConfig) GetURLPath() string {
	return ""
}

func (t *GetKalturaConfig) GetQuery() (string, error) {
	return "", nil
}

func (t *GetKalturaConfig) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetKalturaConfig) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetKalturaConfig) HasErrors() error {
	return nil
}

func (t *GetKalturaConfig) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
