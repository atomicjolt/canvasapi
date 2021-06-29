package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// ListMultipleAssignmentsGradeableStudents A paginated list of students eligible to submit a list of assignments. The caller must have
// permission to view grades for the requested course.
//
// Section-limited instructors will only see students in their own sections.
// https://canvas.instructure.com/doc/api/submissions.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # AssignmentIDs (Optional) Assignments being requested
//
type ListMultipleAssignmentsGradeableStudents struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		AssignmentIDs []string `json:"assignment_ids" url:"assignment_ids,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ListMultipleAssignmentsGradeableStudents) GetMethod() string {
	return "GET"
}

func (t *ListMultipleAssignmentsGradeableStudents) GetURLPath() string {
	path := "courses/{course_id}/assignments/gradeable_students"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListMultipleAssignmentsGradeableStudents) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListMultipleAssignmentsGradeableStudents) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListMultipleAssignmentsGradeableStudents) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListMultipleAssignmentsGradeableStudents) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListMultipleAssignmentsGradeableStudents) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
