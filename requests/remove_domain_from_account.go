package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// RemoveDomainFromAccount Removes an allowed domain from the current account.
// https://canvas.instructure.com/doc/api/content_security_policy_settings.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Query Parameters:
// # Domain (Required) no description
//
type RemoveDomainFromAccount struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Domain string `json:"domain" url:"domain,omitempty"` //  (Required)
	} `json:"query"`
}

func (t *RemoveDomainFromAccount) GetMethod() string {
	return "DELETE"
}

func (t *RemoveDomainFromAccount) GetURLPath() string {
	path := "accounts/{account_id}/csp_settings/domains"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *RemoveDomainFromAccount) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *RemoveDomainFromAccount) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *RemoveDomainFromAccount) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *RemoveDomainFromAccount) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Query.Domain == "" {
		errs = append(errs, "'Domain' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *RemoveDomainFromAccount) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
