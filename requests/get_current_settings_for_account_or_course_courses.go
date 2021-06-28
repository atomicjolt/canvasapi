package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetCurrentSettingsForAccountOrCourseCourses Update multiple modules in an account.
// https://canvas.instructure.com/doc/api/content_security_policy_settings.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type GetCurrentSettingsForAccountOrCourseCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetCurrentSettingsForAccountOrCourseCourses) GetMethod() string {
	return "GET"
}

func (t *GetCurrentSettingsForAccountOrCourseCourses) GetURLPath() string {
	path := "courses/{course_id}/csp_settings"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *GetCurrentSettingsForAccountOrCourseCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *GetCurrentSettingsForAccountOrCourseCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetCurrentSettingsForAccountOrCourseCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetCurrentSettingsForAccountOrCourseCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetCurrentSettingsForAccountOrCourseCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
