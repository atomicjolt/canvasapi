package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetCourseSettings Returns some of a course's settings.
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type GetCourseSettings struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetCourseSettings) GetMethod() string {
	return "GET"
}

func (t *GetCourseSettings) GetURLPath() string {
	path := "courses/{course_id}/settings"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *GetCourseSettings) GetQuery() (string, error) {
	return "", nil
}

func (t *GetCourseSettings) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetCourseSettings) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetCourseSettings) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetCourseSettings) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
