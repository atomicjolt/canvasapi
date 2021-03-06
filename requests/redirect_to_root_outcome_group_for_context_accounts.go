package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// RedirectToRootOutcomeGroupForContextAccounts Convenience redirect to find the root outcome group for a particular
// context. Will redirect to the appropriate outcome group's URL.
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
type RedirectToRootOutcomeGroupForContextAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *RedirectToRootOutcomeGroupForContextAccounts) GetMethod() string {
	return "GET"
}

func (t *RedirectToRootOutcomeGroupForContextAccounts) GetURLPath() string {
	path := "accounts/{account_id}/root_outcome_group"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *RedirectToRootOutcomeGroupForContextAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *RedirectToRootOutcomeGroupForContextAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *RedirectToRootOutcomeGroupForContextAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *RedirectToRootOutcomeGroupForContextAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *RedirectToRootOutcomeGroupForContextAccounts) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
