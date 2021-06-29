package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetVisibleCourseNavigationToolsForSingleCourse Get a list of external tools with the course_navigation placement that have not been hidden in
// course settings and whose visibility settings apply to the requesting user. These tools are the
// same that appear in the course navigation.
//
// The response format is the same as Get visible course navigation tools.
// https://canvas.instructure.com/doc/api/external_tools.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
type GetVisibleCourseNavigationToolsForSingleCourse struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetVisibleCourseNavigationToolsForSingleCourse) GetMethod() string {
	return "GET"
}

func (t *GetVisibleCourseNavigationToolsForSingleCourse) GetURLPath() string {
	path := "courses/{course_id}/external_tools/visible_course_nav_tools"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *GetVisibleCourseNavigationToolsForSingleCourse) GetQuery() (string, error) {
	return "", nil
}

func (t *GetVisibleCourseNavigationToolsForSingleCourse) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetVisibleCourseNavigationToolsForSingleCourse) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetVisibleCourseNavigationToolsForSingleCourse) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetVisibleCourseNavigationToolsForSingleCourse) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
