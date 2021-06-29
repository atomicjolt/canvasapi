package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// Settings Returns all of the settings for the specified account as a JSON object. The caller must be an Account
// admin with the manage_account_settings permission.
// https://canvas.instructure.com/doc/api/accounts.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
type Settings struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *Settings) GetMethod() string {
	return "GET"
}

func (t *Settings) GetURLPath() string {
	path := "accounts/{account_id}/settings"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *Settings) GetQuery() (string, error) {
	return "", nil
}

func (t *Settings) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *Settings) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *Settings) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *Settings) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
