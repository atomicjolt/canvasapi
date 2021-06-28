package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// UploadFileSections Upload a file to a submission.
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
// # SectionID (Required) ID
// # AssignmentID (Required) ID
// # UserID (Required) ID
//
type UploadFileSections struct {
	Path struct {
		SectionID    string `json:"section_id" url:"section_id,omitempty"`       //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
		UserID       string `json:"user_id" url:"user_id,omitempty"`             //  (Required)
	} `json:"path"`
}

func (t *UploadFileSections) GetMethod() string {
	return "POST"
}

func (t *UploadFileSections) GetURLPath() string {
	path := "sections/{section_id}/assignments/{assignment_id}/submissions/{user_id}/files"
	path = strings.ReplaceAll(path, "{section_id}", fmt.Sprintf("%v", t.Path.SectionID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *UploadFileSections) GetQuery() (string, error) {
	return "", nil
}

func (t *UploadFileSections) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *UploadFileSections) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *UploadFileSections) HasErrors() error {
	errs := []string{}
	if t.Path.SectionID == "" {
		errs = append(errs, "'SectionID' is required")
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

func (t *UploadFileSections) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
