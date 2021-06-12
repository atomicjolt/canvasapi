package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// AddMultipleAllowedDomainsToAccount Adds multiple allowed domains for the current account. Note: this will not take effect
// unless CSP is explicitly enabled on this account.
// https://canvas.instructure.com/doc/api/content_security_policy_settings.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Form Parameters:
// # Domains (Required) no description
//
type AddMultipleAllowedDomainsToAccount struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Domains string `json:"domains"` //  (Required)
	} `json:"form"`
}

func (t *AddMultipleAllowedDomainsToAccount) GetMethod() string {
	return "POST"
}

func (t *AddMultipleAllowedDomainsToAccount) GetURLPath() string {
	path := "accounts/{account_id}/csp_settings/domains/batch_create"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *AddMultipleAllowedDomainsToAccount) GetQuery() (string, error) {
	return "", nil
}

func (t *AddMultipleAllowedDomainsToAccount) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *AddMultipleAllowedDomainsToAccount) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Form.Domains == "" {
		errs = append(errs, "'Domains' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *AddMultipleAllowedDomainsToAccount) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}