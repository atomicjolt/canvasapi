package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// CreateJwt Create a unique jwt for using with other canvas services
//
// Generates a different JWT each time it's called, each one expires
// after a short window (1 hour)
// https://canvas.instructure.com/doc/api/jw_ts.html
//
type CreateJwt struct {
}

func (t *CreateJwt) GetMethod() string {
	return "POST"
}

func (t *CreateJwt) GetURLPath() string {
	return ""
}

func (t *CreateJwt) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateJwt) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *CreateJwt) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *CreateJwt) HasErrors() error {
	return nil
}

func (t *CreateJwt) Do(c *canvasapi.Canvas) (*models.JWT, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.JWT{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
