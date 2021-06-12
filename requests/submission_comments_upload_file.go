package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// SubmissionCommentsUploadFile Upload a file to attach to a submission comment
//
// See the {file:file_uploads.html File Upload Documentation} for details on the file upload workflow.
//
// The final step of the file upload workflow will return the attachment data,
// including the new file id. The caller can then PUT the file_id to the
// submission API to attach it to a comment
// https://canvas.instructure.com/doc/api/submission_comments.html
//
// Path Parameters:
// # CourseID (Required) ID
// # AssignmentID (Required) ID
// # UserID (Required) ID
//
type SubmissionCommentsUploadFile struct {
	Path struct {
		CourseID     string `json:"course_id"`     //  (Required)
		AssignmentID string `json:"assignment_id"` //  (Required)
		UserID       string `json:"user_id"`       //  (Required)
	} `json:"path"`
}

func (t *SubmissionCommentsUploadFile) GetMethod() string {
	return "POST"
}

func (t *SubmissionCommentsUploadFile) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/submissions/{user_id}/comments/files"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *SubmissionCommentsUploadFile) GetQuery() (string, error) {
	return "", nil
}

func (t *SubmissionCommentsUploadFile) GetBody() (string, error) {
	return "", nil
}

func (t *SubmissionCommentsUploadFile) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'AssignmentID' is required")
	}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *SubmissionCommentsUploadFile) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
