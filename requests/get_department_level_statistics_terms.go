package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetDepartmentLevelStatisticsTerms Returns numeric statistics about the department and term (or filter).
//
// Shares the same variations on endpoint as the participation data.
// https://canvas.instructure.com/doc/api/analytics.html
//
// Path Parameters:
// # AccountID (Required) ID
// # TermID (Required) ID
//
type GetDepartmentLevelStatisticsTerms struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
		TermID    string `json:"term_id"`    //  (Required)
	} `json:"path"`
}

func (t *GetDepartmentLevelStatisticsTerms) GetMethod() string {
	return "GET"
}

func (t *GetDepartmentLevelStatisticsTerms) GetURLPath() string {
	path := "accounts/{account_id}/analytics/terms/{term_id}/statistics"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{term_id}", fmt.Sprintf("%v", t.Path.TermID))
	return path
}

func (t *GetDepartmentLevelStatisticsTerms) GetQuery() (string, error) {
	return "", nil
}

func (t *GetDepartmentLevelStatisticsTerms) GetBody() (string, error) {
	return "", nil
}

func (t *GetDepartmentLevelStatisticsTerms) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Path.TermID == "" {
		errs = append(errs, "'TermID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetDepartmentLevelStatisticsTerms) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
