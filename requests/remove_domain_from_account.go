package requests

import (
	"fmt"
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
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`

	Query struct {
		Domain string `json:"domain"` //  (Required)
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
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *RemoveDomainFromAccount) GetBody() (string, error) {
	return "", nil
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