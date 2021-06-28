package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// LockOrUnlockCurrentCspSettingsForSubAccountsAndCourses Can only be set if CSP is explicitly enabled or disabled on this account (i.e. "inherited" is false).
// https://canvas.instructure.com/doc/api/content_security_policy_settings.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Form Parameters:
// # SettingsLocked (Required) Whether sub-accounts and courses will be prevented from overriding settings inherited from this account.
//
type LockOrUnlockCurrentCspSettingsForSubAccountsAndCourses struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		SettingsLocked bool `json:"settings_locked" url:"settings_locked,omitempty"` //  (Required)
	} `json:"form"`
}

func (t *LockOrUnlockCurrentCspSettingsForSubAccountsAndCourses) GetMethod() string {
	return "PUT"
}

func (t *LockOrUnlockCurrentCspSettingsForSubAccountsAndCourses) GetURLPath() string {
	path := "accounts/{account_id}/csp_settings/lock"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *LockOrUnlockCurrentCspSettingsForSubAccountsAndCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *LockOrUnlockCurrentCspSettingsForSubAccountsAndCourses) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *LockOrUnlockCurrentCspSettingsForSubAccountsAndCourses) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *LockOrUnlockCurrentCspSettingsForSubAccountsAndCourses) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *LockOrUnlockCurrentCspSettingsForSubAccountsAndCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
