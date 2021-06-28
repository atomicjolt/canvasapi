package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// EditExternalToolCourses Update the specified external tool. Uses same parameters as create
// https://canvas.instructure.com/doc/api/external_tools.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ExternalToolID (Required) ID
//
type EditExternalToolCourses struct {
	Path struct {
		CourseID       string `json:"course_id" url:"course_id,omitempty"`               //  (Required)
		ExternalToolID string `json:"external_tool_id" url:"external_tool_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *EditExternalToolCourses) GetMethod() string {
	return "PUT"
}

func (t *EditExternalToolCourses) GetURLPath() string {
	path := "courses/{course_id}/external_tools/{external_tool_id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{external_tool_id}", fmt.Sprintf("%v", t.Path.ExternalToolID))
	return path
}

func (t *EditExternalToolCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *EditExternalToolCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *EditExternalToolCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *EditExternalToolCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ExternalToolID == "" {
		errs = append(errs, "'ExternalToolID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *EditExternalToolCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
