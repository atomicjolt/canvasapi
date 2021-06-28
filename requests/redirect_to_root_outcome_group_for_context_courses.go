package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// RedirectToRootOutcomeGroupForContextCourses Convenience redirect to find the root outcome group for a particular
// context. Will redirect to the appropriate outcome group's URL.
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type RedirectToRootOutcomeGroupForContextCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *RedirectToRootOutcomeGroupForContextCourses) GetMethod() string {
	return "GET"
}

func (t *RedirectToRootOutcomeGroupForContextCourses) GetURLPath() string {
	path := "courses/{course_id}/root_outcome_group"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *RedirectToRootOutcomeGroupForContextCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *RedirectToRootOutcomeGroupForContextCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *RedirectToRootOutcomeGroupForContextCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *RedirectToRootOutcomeGroupForContextCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *RedirectToRootOutcomeGroupForContextCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
