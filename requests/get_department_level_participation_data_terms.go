package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetDepartmentLevelParticipationDataTerms Returns page view hits summed across all courses in the department. Two
// groupings of these counts are returned; one by day (+by_date+), the other
// by category (+by_category+). The possible categories are announcements,
// assignments, collaborations, conferences, discussions, files, general,
// grades, groups, modules, other, pages, and quizzes.
//
// This and the other department-level endpoints have three variations which
// all return the same style of data but for different subsets of courses. All
// share the prefix /api/v1/accounts/<account_id>/analytics. The possible
// suffixes are:
//
//  * /current: includes all available courses in the default term
//  * /completed: includes all concluded courses in the default term
//  * /terms/<term_id>: includes all available or concluded courses in the
//    given term.
//
// Courses not yet offered or which have been deleted are never included.
//
// /current and /completed are intended for use when the account has only one
// term. /terms/<term_id> is intended for use when the account has multiple
// terms.
//
// The action follows the suffix.
// https://canvas.instructure.com/doc/api/analytics.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
// # Path.TermID (Required) ID
//
type GetDepartmentLevelParticipationDataTerms struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		TermID    string `json:"term_id" url:"term_id,omitempty"`       //  (Required)
	} `json:"path"`
}

func (t *GetDepartmentLevelParticipationDataTerms) GetMethod() string {
	return "GET"
}

func (t *GetDepartmentLevelParticipationDataTerms) GetURLPath() string {
	path := "accounts/{account_id}/analytics/terms/{term_id}/activity"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{term_id}", fmt.Sprintf("%v", t.Path.TermID))
	return path
}

func (t *GetDepartmentLevelParticipationDataTerms) GetQuery() (string, error) {
	return "", nil
}

func (t *GetDepartmentLevelParticipationDataTerms) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetDepartmentLevelParticipationDataTerms) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetDepartmentLevelParticipationDataTerms) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if t.Path.TermID == "" {
		errs = append(errs, "'Path.TermID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetDepartmentLevelParticipationDataTerms) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
