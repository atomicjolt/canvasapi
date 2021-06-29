package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// CourseActivityStreamSummary Returns a summary of the current user's course-specific activity stream.
//
// For full documentation, see the API documentation for the user activity
// stream summary, in the user api.
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
type CourseActivityStreamSummary struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *CourseActivityStreamSummary) GetMethod() string {
	return "GET"
}

func (t *CourseActivityStreamSummary) GetURLPath() string {
	path := "courses/{course_id}/activity_stream/summary"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CourseActivityStreamSummary) GetQuery() (string, error) {
	return "", nil
}

func (t *CourseActivityStreamSummary) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *CourseActivityStreamSummary) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *CourseActivityStreamSummary) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CourseActivityStreamSummary) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
