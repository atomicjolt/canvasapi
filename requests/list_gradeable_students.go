package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListGradeableStudents A paginated list of students eligible to submit the assignment. The caller must have permission to view grades.
//
// If anonymous grading is enabled for the current assignment and the allow_new_anonymous_id parameter is passed,
// the returned data will not include any values identifying the student, but will instead include an
// assignment-specific anonymous ID for each student.
//
// Section-limited instructors will only see students in their own sections.
// https://canvas.instructure.com/doc/api/submissions.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.AssignmentID (Required) ID
//
type ListGradeableStudents struct {
	Path struct {
		CourseID     string `json:"course_id" url:"course_id,omitempty"`         //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListGradeableStudents) GetMethod() string {
	return "GET"
}

func (t *ListGradeableStudents) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/gradeable_students"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	return path
}

func (t *ListGradeableStudents) GetQuery() (string, error) {
	return "", nil
}

func (t *ListGradeableStudents) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListGradeableStudents) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListGradeableStudents) HasErrors() error {
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

func (t *ListGradeableStudents) Do(c *canvasapi.Canvas) ([]*models.UserDisplay, *canvasapi.PagedResource, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.UserDisplay{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, nil, err
	}

	pagedResource, err := canvasapi.ExtractPagedResource(response.Header)
	if err != nil {
		return nil, nil, err
	}

	return ret, pagedResource, nil
}
