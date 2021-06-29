package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetCurrentSettingsForAccountOrCourseAccounts Update multiple modules in an account.
// https://canvas.instructure.com/doc/api/content_security_policy_settings.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
type GetCurrentSettingsForAccountOrCourseAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetCurrentSettingsForAccountOrCourseAccounts) GetMethod() string {
	return "GET"
}

func (t *GetCurrentSettingsForAccountOrCourseAccounts) GetURLPath() string {
	path := "accounts/{account_id}/csp_settings"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *GetCurrentSettingsForAccountOrCourseAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *GetCurrentSettingsForAccountOrCourseAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetCurrentSettingsForAccountOrCourseAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetCurrentSettingsForAccountOrCourseAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetCurrentSettingsForAccountOrCourseAccounts) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
