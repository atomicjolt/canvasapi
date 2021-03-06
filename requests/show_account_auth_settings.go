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

// ShowAccountAuthSettings The way to get the current state of each account level setting
// that's relevant to Single Sign On configuration
//
// You can list the current state of each setting with "update_sso_settings"
// https://canvas.instructure.com/doc/api/authentication_providers.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
type ShowAccountAuthSettings struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ShowAccountAuthSettings) GetMethod() string {
	return "GET"
}

func (t *ShowAccountAuthSettings) GetURLPath() string {
	path := "accounts/{account_id}/sso_settings"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ShowAccountAuthSettings) GetQuery() (string, error) {
	return "", nil
}

func (t *ShowAccountAuthSettings) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ShowAccountAuthSettings) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ShowAccountAuthSettings) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowAccountAuthSettings) Do(c *canvasapi.Canvas) (*models.SSOSettings, error) {
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
