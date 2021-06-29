package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// SelectProvisionalGrade Choose which provisional grade the student should receive for a submission.
// The caller must be the final grader for the assignment or an admin with :select_final_grade rights.
// https://canvas.instructure.com/doc/api/moderated_grading.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.AssignmentID (Required) ID
// # Path.ProvisionalGradeID (Required) ID
//
type SelectProvisionalGrade struct {
	Path struct {
		CourseID           string `json:"course_id" url:"course_id,omitempty"`                       //  (Required)
		AssignmentID       string `json:"assignment_id" url:"assignment_id,omitempty"`               //  (Required)
		ProvisionalGradeID string `json:"provisional_grade_id" url:"provisional_grade_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *SelectProvisionalGrade) GetMethod() string {
	return "PUT"
}

func (t *SelectProvisionalGrade) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/provisional_grades/{provisional_grade_id}/select"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{provisional_grade_id}", fmt.Sprintf("%v", t.Path.ProvisionalGradeID))
	return path
}

func (t *SelectProvisionalGrade) GetQuery() (string, error) {
	return "", nil
}

func (t *SelectProvisionalGrade) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *SelectProvisionalGrade) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *SelectProvisionalGrade) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'Path.AssignmentID' is required")
	}
	if t.Path.ProvisionalGradeID == "" {
		errs = append(errs, "'Path.ProvisionalGradeID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *SelectProvisionalGrade) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
