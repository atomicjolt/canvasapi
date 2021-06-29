package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// MarkSubmissionAsReadSections No request fields are necessary.
//
// On success, the response will be 204 No Content with an empty body.
// https://canvas.instructure.com/doc/api/submissions.html
//
// Path Parameters:
// # Path.SectionID (Required) ID
// # Path.AssignmentID (Required) ID
// # Path.UserID (Required) ID
//
type MarkSubmissionAsReadSections struct {
	Path struct {
		SectionID    string `json:"section_id" url:"section_id,omitempty"`       //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
		UserID       string `json:"user_id" url:"user_id,omitempty"`             //  (Required)
	} `json:"path"`
}

func (t *MarkSubmissionAsReadSections) GetMethod() string {
	return "PUT"
}

func (t *MarkSubmissionAsReadSections) GetURLPath() string {
	path := "sections/{section_id}/assignments/{assignment_id}/submissions/{user_id}/read"
	path = strings.ReplaceAll(path, "{section_id}", fmt.Sprintf("%v", t.Path.SectionID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *MarkSubmissionAsReadSections) GetQuery() (string, error) {
	return "", nil
}

func (t *MarkSubmissionAsReadSections) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *MarkSubmissionAsReadSections) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *MarkSubmissionAsReadSections) HasErrors() error {
	errs := []string{}
	if t.Path.SectionID == "" {
		errs = append(errs, "'Path.SectionID' is required")
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

func (t *MarkSubmissionAsReadSections) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
