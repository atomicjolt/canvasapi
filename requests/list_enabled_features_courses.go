package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ListEnabledFeaturesCourses A paginated list of all features that are enabled on a given Account, Course, or User.
// Only the feature names are returned.
// https://canvas.instructure.com/doc/api/feature_flags.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type ListEnabledFeaturesCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListEnabledFeaturesCourses) GetMethod() string {
	return "GET"
}

func (t *ListEnabledFeaturesCourses) GetURLPath() string {
	path := "courses/{course_id}/features/enabled"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListEnabledFeaturesCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ListEnabledFeaturesCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListEnabledFeaturesCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListEnabledFeaturesCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListEnabledFeaturesCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
