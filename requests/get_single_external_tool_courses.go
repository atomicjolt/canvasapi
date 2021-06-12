package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetSingleExternalToolCourses Returns the specified external tool.
// https://canvas.instructure.com/doc/api/external_tools.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ExternalToolID (Required) ID
//
type GetSingleExternalToolCourses struct {
	Path struct {
		CourseID       string `json:"course_id"`        //  (Required)
		ExternalToolID string `json:"external_tool_id"` //  (Required)
	} `json:"path"`
}

func (t *GetSingleExternalToolCourses) GetMethod() string {
	return "GET"
}

func (t *GetSingleExternalToolCourses) GetURLPath() string {
	path := "courses/{course_id}/external_tools/{external_tool_id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{external_tool_id}", fmt.Sprintf("%v", t.Path.ExternalToolID))
	return path
}

func (t *GetSingleExternalToolCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *GetSingleExternalToolCourses) GetBody() (string, error) {
	return "", nil
}

func (t *GetSingleExternalToolCourses) HasErrors() error {
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

func (t *GetSingleExternalToolCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}