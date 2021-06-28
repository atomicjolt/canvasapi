package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// CourseTodoItems Returns the current user's course-specific todo items.
//
// For full documentation, see the API documentation for the user todo items, in the user api.
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type CourseTodoItems struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *CourseTodoItems) GetMethod() string {
	return "GET"
}

func (t *CourseTodoItems) GetURLPath() string {
	path := "courses/{course_id}/todo"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CourseTodoItems) GetQuery() (string, error) {
	return "", nil
}

func (t *CourseTodoItems) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *CourseTodoItems) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *CourseTodoItems) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CourseTodoItems) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
