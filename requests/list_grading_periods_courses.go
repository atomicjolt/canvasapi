package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ListGradingPeriodsCourses Returns the paginated list of grading periods for the current course.
// https://canvas.instructure.com/doc/api/grading_periods.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type ListGradingPeriodsCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListGradingPeriodsCourses) GetMethod() string {
	return "GET"
}

func (t *ListGradingPeriodsCourses) GetURLPath() string {
	path := "courses/{course_id}/grading_periods"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListGradingPeriodsCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ListGradingPeriodsCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListGradingPeriodsCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListGradingPeriodsCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListGradingPeriodsCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
