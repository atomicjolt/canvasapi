package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// RetrieveReportedCspViolationsForAccount Must be called on a root account.
// https://canvas.instructure.com/doc/api/content_security_policy_settings.html
//
// Path Parameters:
// # AccountID (Required) ID
//
type RetrieveReportedCspViolationsForAccount struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`
}

func (t *RetrieveReportedCspViolationsForAccount) GetMethod() string {
	return "GET"
}

func (t *RetrieveReportedCspViolationsForAccount) GetURLPath() string {
	path := "accounts/{account_id}/csp_log"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *RetrieveReportedCspViolationsForAccount) GetQuery() (string, error) {
	return "", nil
}

func (t *RetrieveReportedCspViolationsForAccount) GetBody() (string, error) {
	return "", nil
}

func (t *RetrieveReportedCspViolationsForAccount) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *RetrieveReportedCspViolationsForAccount) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
