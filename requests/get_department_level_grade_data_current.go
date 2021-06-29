package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetDepartmentLevelGradeDataCurrent Returns the distribution of grades for students in courses in the
// department.  Each data point is one student's current grade in one course;
// if a student is in multiple courses, he contributes one value per course,
// but if he's enrolled multiple times in the same course (e.g. a lecture
// section and a lab section), he only constributes on value for that course.
//
// Grades are binned to the nearest integer score; anomalous grades outside
// the 0 to 100 range are ignored. The raw counts are returned, not yet
// normalized by the total count.
//
// Shares the same variations on endpoint as the participation data.
// https://canvas.instructure.com/doc/api/analytics.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
type GetDepartmentLevelGradeDataCurrent struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetDepartmentLevelGradeDataCurrent) GetMethod() string {
	return "GET"
}

func (t *GetDepartmentLevelGradeDataCurrent) GetURLPath() string {
	path := "accounts/{account_id}/analytics/current/grades"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *GetDepartmentLevelGradeDataCurrent) GetQuery() (string, error) {
	return "", nil
}

func (t *GetDepartmentLevelGradeDataCurrent) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetDepartmentLevelGradeDataCurrent) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetDepartmentLevelGradeDataCurrent) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetDepartmentLevelGradeDataCurrent) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
