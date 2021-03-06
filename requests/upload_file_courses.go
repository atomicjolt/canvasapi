package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// UploadFileCourses Upload a file to a submission.
//
// This API endpoint is the first step in uploading a file to a submission as a student.
// See the {file:file_uploads.html File Upload Documentation} for details on the file upload workflow.
//
// The final step of the file upload workflow will return the attachment data,
// including the new file id. The caller can then POST to submit the
// +online_upload+ assignment with these file ids.
// https://canvas.instructure.com/doc/api/submissions.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.AssignmentID (Required) ID
// # Path.UserID (Required) ID
//
type UploadFileCourses struct {
	Path struct {
		CourseID     string `json:"course_id" url:"course_id,omitempty"`         //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
		UserID       string `json:"user_id" url:"user_id,omitempty"`             //  (Required)
	} `json:"path"`
}

func (t *UploadFileCourses) GetMethod() string {
	return "POST"
}

func (t *UploadFileCourses) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/submissions/{user_id}/files"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *UploadFileCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *UploadFileCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *UploadFileCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *UploadFileCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'Path.AssignmentID' is required")
	}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UploadFileCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
