package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// CourseActivityStream Returns the current user's course-specific activity stream, paginated.
//
// For full documentation, see the API documentation for the user activity
// stream, in the user api.
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type CourseActivityStream struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`
}

func (t *CourseActivityStream) GetMethod() string {
	return "GET"
}

func (t *CourseActivityStream) GetURLPath() string {
	path := "courses/{course_id}/activity_stream"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CourseActivityStream) GetQuery() (string, error) {
	return "", nil
}

func (t *CourseActivityStream) GetBody() (string, error) {
	return "", nil
}

func (t *CourseActivityStream) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CourseActivityStream) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
