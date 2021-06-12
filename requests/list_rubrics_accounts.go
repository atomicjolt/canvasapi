package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ListRubricsAccounts Returns the paginated list of active rubrics for the current context.
// https://canvas.instructure.com/doc/api/rubrics.html
//
// Path Parameters:
// # AccountID (Required) ID
//
type ListRubricsAccounts struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`
}

func (t *ListRubricsAccounts) GetMethod() string {
	return "GET"
}

func (t *ListRubricsAccounts) GetURLPath() string {
	path := "accounts/{account_id}/rubrics"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ListRubricsAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *ListRubricsAccounts) GetBody() (string, error) {
	return "", nil
}

func (t *ListRubricsAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListRubricsAccounts) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
