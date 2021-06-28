package requests

import (
	"net/url"

	"github.com/atomicjolt/canvasapi"
)

// ListEnvironmentFeatures Return a hash of global feature settings that pertain to the
// Canvas user interface. This is the same information supplied to the
// web interface as +ENV.FEATURES+.
// https://canvas.instructure.com/doc/api/feature_flags.html
//
type ListEnvironmentFeatures struct {
}

func (t *ListEnvironmentFeatures) GetMethod() string {
	return "GET"
}

func (t *ListEnvironmentFeatures) GetURLPath() string {
	return ""
}

func (t *ListEnvironmentFeatures) GetQuery() (string, error) {
	return "", nil
}

func (t *ListEnvironmentFeatures) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListEnvironmentFeatures) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListEnvironmentFeatures) HasErrors() error {
	return nil
}

func (t *ListEnvironmentFeatures) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
