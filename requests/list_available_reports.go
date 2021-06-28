package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ListAvailableReports Returns a paginated list of reports for the current context.
// https://canvas.instructure.com/doc/api/account_reports.html
//
// Path Parameters:
// # AccountID (Required) ID
//
type ListAvailableReports struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListAvailableReports) GetMethod() string {
	return "GET"
}

func (t *ListAvailableReports) GetURLPath() string {
	path := "accounts/{account_id}/reports"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ListAvailableReports) GetQuery() (string, error) {
	return "", nil
}

func (t *ListAvailableReports) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListAvailableReports) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListAvailableReports) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListAvailableReports) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
