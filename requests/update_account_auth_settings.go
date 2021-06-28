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

// UpdateAccountAuthSettings For various cases of mixed SSO configurations, you may need to set some
// configuration at the account level to handle the particulars of your
// setup.
//
// This endpoint accepts a PUT request to set several possible account
// settings. All setting are optional on each request, any that are not
// provided at all are simply retained as is.  Any that provide the key but
// a null-ish value (blank string, null, undefined) will be UN-set.
//
// You can list the current state of each setting with "show_sso_settings"
// https://canvas.instructure.com/doc/api/authentication_providers.html
//
// Path Parameters:
// # AccountID (Required) ID
//
type UpdateAccountAuthSettings struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *UpdateAccountAuthSettings) GetMethod() string {
	return "PUT"
}

func (t *UpdateAccountAuthSettings) GetURLPath() string {
	path := "accounts/{account_id}/sso_settings"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *UpdateAccountAuthSettings) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateAccountAuthSettings) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *UpdateAccountAuthSettings) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *UpdateAccountAuthSettings) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateAccountAuthSettings) Do(c *canvasapi.Canvas) (*models.SSOSettings, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.SSOSettings{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
