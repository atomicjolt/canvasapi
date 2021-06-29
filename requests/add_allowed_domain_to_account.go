package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// AddAllowedDomainToAccount Adds an allowed domain for the current account. Note: this will not take effect
// unless CSP is explicitly enabled on this account.
// https://canvas.instructure.com/doc/api/content_security_policy_settings.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
// Form Parameters:
// # Form.Domain (Required) no description
//
type AddAllowedDomainToAccount struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Domain string `json:"domain" url:"domain,omitempty"` //  (Required)
	} `json:"form"`
}

func (t *AddAllowedDomainToAccount) GetMethod() string {
	return "POST"
}

func (t *AddAllowedDomainToAccount) GetURLPath() string {
	path := "accounts/{account_id}/csp_settings/domains"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *AddAllowedDomainToAccount) GetQuery() (string, error) {
	return "", nil
}

func (t *AddAllowedDomainToAccount) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *AddAllowedDomainToAccount) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *AddAllowedDomainToAccount) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if t.Form.Domain == "" {
		errs = append(errs, "'Form.Domain' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *AddAllowedDomainToAccount) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
