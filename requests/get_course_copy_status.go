package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetCourseCopyStatus DEPRECATED: Please use the {api:ContentMigrationsController#create Content Migrations API}
//
// Retrieve the status of a course copy
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
type GetCourseCopyStatus struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`
}

func (t *GetCourseCopyStatus) GetMethod() string {
	return "GET"
}

func (t *GetCourseCopyStatus) GetURLPath() string {
	path := "courses/{course_id}/course_copy/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetCourseCopyStatus) GetQuery() (string, error) {
	return "", nil
}

func (t *GetCourseCopyStatus) GetBody() (string, error) {
	return "", nil
}

func (t *GetCourseCopyStatus) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetCourseCopyStatus) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
