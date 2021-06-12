package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// DeleteGradingPeriodCourses <b>204 No Content</b> response code is returned if the deletion was
// successful.
// https://canvas.instructure.com/doc/api/grading_periods.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
type DeleteGradingPeriodCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
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

func (t *DeleteGradingPeriodCourses) GetBody() (string, error) {
	return "", nil
}

func (t *DeleteGradingPeriodCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
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