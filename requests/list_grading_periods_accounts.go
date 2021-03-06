package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ListGradingPeriodsAccounts Returns the paginated list of grading periods for the current course.
// https://canvas.instructure.com/doc/api/grading_periods.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
type ListGradingPeriodsAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListGradingPeriodsAccounts) GetMethod() string {
	return "GET"
}

func (t *ListGradingPeriodsAccounts) GetURLPath() string {
	path := "accounts/{account_id}/grading_periods"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ListGradingPeriodsAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *ListGradingPeriodsAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListGradingPeriodsAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListGradingPeriodsAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListGradingPeriodsAccounts) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
