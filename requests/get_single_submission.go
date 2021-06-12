package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetSingleSubmission Get a single submission, based on submission id.
// https://canvas.instructure.com/doc/api/plagiarism_detection_submissions.html
//
// Path Parameters:
// # AssignmentID (Required) ID
// # SubmissionID (Required) ID
//
type GetSingleSubmission struct {
	Path struct {
		AssignmentID string `json:"assignment_id"` //  (Required)
		SubmissionID string `json:"submission_id"` //  (Required)
	} `json:"path"`
}

func (t *GetSingleSubmission) GetMethod() string {
	return "GET"
}

func (t *GetSingleSubmission) GetURLPath() string {
	path := "/lti/assignments/{assignment_id}/submissions/{submission_id}"
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{submission_id}", fmt.Sprintf("%v", t.Path.SubmissionID))
	return path
}

func (t *GetSingleSubmission) GetQuery() (string, error) {
	return "", nil
}

func (t *GetSingleSubmission) GetBody() (string, error) {
	return "", nil
}

func (t *GetSingleSubmission) HasErrors() error {
	errs := []string{}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'AssignmentID' is required")
	}
	if t.Path.SubmissionID == "" {
		errs = append(errs, "'SubmissionID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleSubmission) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
