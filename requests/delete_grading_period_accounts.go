package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// DeleteGradingPeriodAccounts <b>204 No Content</b> response code is returned if the deletion was
// successful.
// https://canvas.instructure.com/doc/api/grading_periods.html
//
// Path Parameters:
// # AccountID (Required) ID
// # ID (Required) ID
//
type DeleteGradingPeriodAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`
}

func (t *DeleteGradingPeriodAccounts) GetMethod() string {
	return "DELETE"
}

func (t *DeleteGradingPeriodAccounts) GetURLPath() string {
	path := "accounts/{account_id}/grading_periods/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteGradingPeriodAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteGradingPeriodAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteGradingPeriodAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteGradingPeriodAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteGradingPeriodAccounts) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
