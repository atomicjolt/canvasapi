package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetDepartmentLevelStatisticsCompleted Returns numeric statistics about the department and term (or filter).
//
// Shares the same variations on endpoint as the participation data.
// https://canvas.instructure.com/doc/api/analytics.html
//
// Path Parameters:
// # AccountID (Required) ID
//
type GetDepartmentLevelStatisticsCompleted struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`
}

func (t *GetDepartmentLevelStatisticsCompleted) GetMethod() string {
	return "GET"
}

func (t *GetDepartmentLevelStatisticsCompleted) GetURLPath() string {
	path := "accounts/{account_id}/analytics/completed/statistics"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *GetDepartmentLevelStatisticsCompleted) GetQuery() (string, error) {
	return "", nil
}

func (t *GetDepartmentLevelStatisticsCompleted) GetBody() (string, error) {
	return "", nil
}

func (t *GetDepartmentLevelStatisticsCompleted) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetDepartmentLevelStatisticsCompleted) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
