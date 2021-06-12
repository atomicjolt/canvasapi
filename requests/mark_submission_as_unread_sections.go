package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// MarkSubmissionAsUnreadSections No request fields are necessary.
//
// On success, the response will be 204 No Content with an empty body.
// https://canvas.instructure.com/doc/api/submissions.html
//
// Path Parameters:
// # SectionID (Required) ID
// # AssignmentID (Required) ID
// # UserID (Required) ID
//
type MarkSubmissionAsUnreadSections struct {
	Path struct {
		SectionID    string `json:"section_id"`    //  (Required)
		AssignmentID string `json:"assignment_id"` //  (Required)
		UserID       string `json:"user_id"`       //  (Required)
	} `json:"path"`
}

func (t *MarkSubmissionAsUnreadSections) GetMethod() string {
	return "DELETE"
}

func (t *MarkSubmissionAsUnreadSections) GetURLPath() string {
	path := "sections/{section_id}/assignments/{assignment_id}/submissions/{user_id}/read"
	path = strings.ReplaceAll(path, "{section_id}", fmt.Sprintf("%v", t.Path.SectionID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *MarkSubmissionAsUnreadSections) GetQuery() (string, error) {
	return "", nil
}

func (t *MarkSubmissionAsUnreadSections) GetBody() (string, error) {
	return "", nil
}

func (t *MarkSubmissionAsUnreadSections) HasErrors() error {
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

func (t *MarkSubmissionAsUnreadSections) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}