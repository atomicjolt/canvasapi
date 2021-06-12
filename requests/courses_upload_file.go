package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// CoursesUploadFile Upload a file to the course.
//
// This API endpoint is the first step in uploading a file to a course.
// See the {file:file_uploads.html File Upload Documentation} for details on
// the file upload workflow.
//
// Only those with the "Manage Files" permission on a course can upload files
// to the course. By default, this is Teachers, TAs and Designers.
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type CoursesUploadFile struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`
}

func (t *CoursesUploadFile) GetMethod() string {
	return "POST"
}

func (t *CoursesUploadFile) GetURLPath() string {
	path := "courses/{course_id}/files"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CoursesUploadFile) GetQuery() (string, error) {
	return "", nil
}

func (t *CoursesUploadFile) GetBody() (string, error) {
	return "", nil
}

func (t *CoursesUploadFile) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CoursesUploadFile) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
