package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetQuotaInformationCourses Returns the total and used storage quota for the course, group, or user.
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type GetQuotaInformationCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`
}

func (t *GetQuotaInformationCourses) GetMethod() string {
	return "GET"
}

func (t *GetQuotaInformationCourses) GetURLPath() string {
	path := "courses/{course_id}/files/quota"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *GetQuotaInformationCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *GetQuotaInformationCourses) GetBody() (string, error) {
	return "", nil
}

func (t *GetQuotaInformationCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetQuotaInformationCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
