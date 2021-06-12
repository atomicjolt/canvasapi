package requests

import (
	"fmt"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// GetPandataEventsJwtTokenAndItsExpirationDate Returns a jwt auth and props token that can be used to send events to
// Pandata.
//
// NOTE: This is currently only available to the mobile developer keys.
// https://canvas.instructure.com/doc/api/users.html
//
// Form Parameters:
// # AppKey (Optional) The pandata events appKey for this mobile app
//
type GetPandataEventsJwtTokenAndItsExpirationDate struct {
	Form struct {
		AppKey string `json:"app_key"` //  (Optional)
	} `json:"form"`
}

func (t *GetPandataEventsJwtTokenAndItsExpirationDate) GetMethod() string {
	return "POST"
}

func (t *GetPandataEventsJwtTokenAndItsExpirationDate) GetURLPath() string {
	return ""
}

func (t *GetPandataEventsJwtTokenAndItsExpirationDate) GetQuery() (string, error) {
	return "", nil
}

func (t *GetPandataEventsJwtTokenAndItsExpirationDate) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *GetPandataEventsJwtTokenAndItsExpirationDate) HasErrors() error {
	return nil
}

func (t *GetPandataEventsJwtTokenAndItsExpirationDate) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
