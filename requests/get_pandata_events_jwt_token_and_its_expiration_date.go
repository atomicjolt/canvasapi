package requests

import (
	"encoding/json"
	"net/url"

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
		AppKey string `json:"app_key" url:"app_key,omitempty"` //  (Optional)
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

func (t *GetPandataEventsJwtTokenAndItsExpirationDate) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *GetPandataEventsJwtTokenAndItsExpirationDate) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
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
