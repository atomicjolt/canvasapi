package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// DeleteGradingPeriodCourses <b>204 No Content</b> response code is returned if the deletion was
// successful.
// https://canvas.instructure.com/doc/api/grading_periods.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.ID (Required) ID
//
type DeleteGradingPeriodCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`
}

func (t *DeleteGradingPeriodCourses) GetMethod() string {
	return "DELETE"
}

func (t *DeleteGradingPeriodCourses) GetURLPath() string {
	path := "courses/{course_id}/grading_periods/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteGradingPeriodCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteGradingPeriodCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteGradingPeriodCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteGradingPeriodCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteGradingPeriodCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
