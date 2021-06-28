package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// UpdatePublicJwk Rotate the public key in jwk format when using lti services
// https://canvas.instructure.com/doc/api/public_jwk.html
//
// Form Parameters:
// # PublicJwk (Required) The new public jwk that will be set to the tools current public jwk.
//
type UpdatePublicJwk struct {
	Form struct {
		PublicJwk string `json:"public_jwk" url:"public_jwk,omitempty"` //  (Required)
	} `json:"form"`
}

func (t *UpdatePublicJwk) GetMethod() string {
	return "PUT"
}

func (t *UpdatePublicJwk) GetURLPath() string {
	return ""
}

func (t *UpdatePublicJwk) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdatePublicJwk) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdatePublicJwk) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdatePublicJwk) HasErrors() error {
	errs := []string{}
	if t.Form.PublicJwk == "" {
		errs = append(errs, "'PublicJwk' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdatePublicJwk) Do(c *canvasapi.Canvas) (*models.DeveloperKey, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.DeveloperKey{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
