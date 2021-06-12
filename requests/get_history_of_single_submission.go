package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetHistoryOfSingleSubmission Get a list of all attempts made for a submission, based on submission id.
// https://canvas.instructure.com/doc/api/plagiarism_detection_submissions.html
//
// Path Parameters:
// # AssignmentID (Required) ID
// # SubmissionID (Required) ID
//
type GetHistoryOfSingleSubmission struct {
	Path struct {
		AssignmentID string `json:"assignment_id"` //  (Required)
		SubmissionID string `json:"submission_id"` //  (Required)
	} `json:"path"`
}

func (t *GetHistoryOfSingleSubmission) GetMethod() string {
	return "GET"
}

func (t *GetHistoryOfSingleSubmission) GetURLPath() string {
	path := "/lti/assignments/{assignment_id}/submissions/{submission_id}/history"
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{submission_id}", fmt.Sprintf("%v", t.Path.SubmissionID))
	return path
}

func (t *GetHistoryOfSingleSubmission) GetQuery() (string, error) {
	return "", nil
}

func (t *GetHistoryOfSingleSubmission) GetBody() (string, error) {
	return "", nil
}

func (t *GetHistoryOfSingleSubmission) HasErrors() error {
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

func (t *GetHistoryOfSingleSubmission) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
