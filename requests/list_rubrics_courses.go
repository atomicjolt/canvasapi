package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ListRubricsCourses Returns the paginated list of active rubrics for the current context.
// https://canvas.instructure.com/doc/api/rubrics.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
type ListRubricsCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListRubricsCourses) GetMethod() string {
	return "GET"
}

func (t *ListRubricsCourses) GetURLPath() string {
	path := "courses/{course_id}/rubrics"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListRubricsCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ListRubricsCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListRubricsCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListRubricsCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListRubricsCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
