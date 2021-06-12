package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// UpdateAuthenticationProvider Update an authentication provider using the same options as the create endpoint.
// You can not update an existing provider to a new authentication type.
// https://canvas.instructure.com/doc/api/authentication_providers.html
//
// Path Parameters:
// # AccountID (Required) ID
// # ID (Required) ID
//
type UpdateAuthenticationProvider struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
		ID        string `json:"id"`         //  (Required)
	} `json:"path"`
}

func (t *UpdateAuthenticationProvider) GetMethod() string {
	return "PUT"
}

func (t *UpdateAuthenticationProvider) GetURLPath() string {
	path := "accounts/{account_id}/authentication_providers/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateAuthenticationProvider) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateAuthenticationProvider) GetBody() (string, error) {
	return "", nil
}

func (t *UpdateAuthenticationProvider) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateAuthenticationProvider) Do(c *canvasapi.Canvas) (*models.AuthenticationProvider, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.AuthenticationProvider{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
