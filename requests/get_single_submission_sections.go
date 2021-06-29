package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// GetSingleSubmissionSections Get a single submission, based on user id.
// https://canvas.instructure.com/doc/api/submissions.html
//
// Path Parameters:
// # SectionID (Required) ID
// # AssignmentID (Required) ID
// # UserID (Required) ID
//
// Query Parameters:
// # Include (Optional) . Must be one of submission_history, submission_comments, rubric_assessment, full_rubric_assessment, visibility, course, user, read_statusAssociations to include with the group.
//
type GetSingleSubmissionSections struct {
	Path struct {
		SectionID    string `json:"section_id" url:"section_id,omitempty"`       //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
		UserID       string `json:"user_id" url:"user_id,omitempty"`             //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of submission_history, submission_comments, rubric_assessment, full_rubric_assessment, visibility, course, user, read_status
	} `json:"query"`
}

func (t *GetSingleSubmissionSections) GetMethod() string {
	return "GET"
}

func (t *GetSingleSubmissionSections) GetURLPath() string {
	path := "sections/{section_id}/assignments/{assignment_id}/submissions/{user_id}"
	path = strings.ReplaceAll(path, "{section_id}", fmt.Sprintf("%v", t.Path.SectionID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *GetSingleSubmissionSections) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *GetSingleSubmissionSections) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSingleSubmissionSections) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSingleSubmissionSections) HasErrors() error {
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
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"submission_history", "submission_comments", "rubric_assessment", "full_rubric_assessment", "visibility", "course", "user", "read_status"}, v) {
			errs = append(errs, "Include must be one of submission_history, submission_comments, rubric_assessment, full_rubric_assessment, visibility, course, user, read_status")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleSubmissionSections) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
