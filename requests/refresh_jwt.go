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

// RefreshJwt Refresh a JWT for use with other canvas services
//
// Generates a different JWT each time it's called, each one expires
// after a short window (1 hour).
// https://canvas.instructure.com/doc/api/jw_ts.html
//
// Form Parameters:
// # Jwt (Required) An existing JWT token to be refreshed. The new token will have
//    the same context and workflows as the existing token.
//
type RefreshJwt struct {
	Form struct {
		Jwt string `json:"jwt" url:"jwt,omitempty"` //  (Required)
	} `json:"form"`
}

func (t *RefreshJwt) GetMethod() string {
	return "POST"
}

func (t *RefreshJwt) GetURLPath() string {
	return ""
}

func (t *RefreshJwt) GetQuery() (string, error) {
	return "", nil
}

func (t *RefreshJwt) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *RefreshJwt) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *RefreshJwt) HasErrors() error {
	errs := []string{}
	if t.Form.Jwt == "" {
		errs = append(errs, "'Jwt' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *RefreshJwt) Do(c *canvasapi.Canvas) (*models.JWT, error) {
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
