package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// BulkSelectProvisionalGrades Choose which provisional grades will be received by associated students for an assignment.
// The caller must be the final grader for the assignment or an admin with :select_final_grade rights.
// https://canvas.instructure.com/doc/api/moderated_grading.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.AssignmentID (Required) ID
//
type BulkSelectProvisionalGrades struct {
	Path struct {
		CourseID     string `json:"course_id" url:"course_id,omitempty"`         //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *BulkSelectProvisionalGrades) GetMethod() string {
	return "PUT"
}

func (t *BulkSelectProvisionalGrades) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/provisional_grades/bulk_select"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	return path
}

func (t *BulkSelectProvisionalGrades) GetQuery() (string, error) {
	return "", nil
}

func (t *BulkSelectProvisionalGrades) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *BulkSelectProvisionalGrades) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *BulkSelectProvisionalGrades) HasErrors() error {
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

func (t *BulkSelectProvisionalGrades) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
