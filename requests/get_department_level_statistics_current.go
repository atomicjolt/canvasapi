package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetDepartmentLevelStatisticsCurrent Returns numeric statistics about the department and term (or filter).
//
// Shares the same variations on endpoint as the participation data.
// https://canvas.instructure.com/doc/api/analytics.html
//
// Path Parameters:
// # AccountID (Required) ID
//
type GetDepartmentLevelStatisticsCurrent struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`
}

func (t *GetDepartmentLevelStatisticsCurrent) GetMethod() string {
	return "GET"
}

func (t *GetDepartmentLevelStatisticsCurrent) GetURLPath() string {
	path := "accounts/{account_id}/analytics/current/statistics"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *GetDepartmentLevelStatisticsCurrent) GetQuery() (string, error) {
	return "", nil
}

func (t *GetDepartmentLevelStatisticsCurrent) GetBody() (string, error) {
	return "", nil
}

func (t *GetDepartmentLevelStatisticsCurrent) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetDepartmentLevelStatisticsCurrent) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
