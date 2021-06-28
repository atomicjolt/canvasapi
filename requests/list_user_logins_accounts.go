package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ListUserLoginsAccounts Given a user ID, return a paginated list of that user's logins for the given account.
// https://canvas.instructure.com/doc/api/logins.html
//
// Path Parameters:
// # AccountID (Required) ID
//
type ListUserLoginsAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListUserLoginsAccounts) GetMethod() string {
	return "GET"
}

func (t *ListUserLoginsAccounts) GetURLPath() string {
	path := "accounts/{account_id}/logins"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ListUserLoginsAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *ListUserLoginsAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListUserLoginsAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListUserLoginsAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListUserLoginsAccounts) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
