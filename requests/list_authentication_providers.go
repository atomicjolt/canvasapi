package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListAuthenticationProviders Returns a paginated list of authentication providers
// https://canvas.instructure.com/doc/api/authentication_providers.html
//
// Path Parameters:
// # AccountID (Required) ID
//
type ListAuthenticationProviders struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListAuthenticationProviders) GetMethod() string {
	return "GET"
}

func (t *ListAuthenticationProviders) GetURLPath() string {
	path := "accounts/{account_id}/authentication_providers"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ListAuthenticationProviders) GetQuery() (string, error) {
	return "", nil
}

func (t *ListAuthenticationProviders) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListAuthenticationProviders) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListAuthenticationProviders) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListAuthenticationProviders) Do(c *canvasapi.Canvas) ([]*models.AuthenticationProvider, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.AuthenticationProvider{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
