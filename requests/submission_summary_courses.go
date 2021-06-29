package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// SubmissionSummaryCourses Returns the number of submissions for the given assignment based on gradeable students
// that fall into three categories: graded, ungraded, not submitted.
// https://canvas.instructure.com/doc/api/submissions.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.AssignmentID (Required) ID
//
// Query Parameters:
// # Query.Grouped (Optional) If this argument is true, the response will take into account student groups.
//
type SubmissionSummaryCourses struct {
	Path struct {
		CourseID     string `json:"course_id" url:"course_id,omitempty"`         //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Grouped bool `json:"grouped" url:"grouped,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *SubmissionSummaryCourses) GetMethod() string {
	return "GET"
}

func (t *SubmissionSummaryCourses) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/submission_summary"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	return path
}

func (t *SubmissionSummaryCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *SubmissionSummaryCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *SubmissionSummaryCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *SubmissionSummaryCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'Path.AssignmentID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *SubmissionSummaryCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
