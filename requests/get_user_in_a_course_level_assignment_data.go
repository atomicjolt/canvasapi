package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetUserInACourseLevelAssignmentData Returns a list of assignments for the course sorted by due date. For
// each assignment returns basic assignment information, the grade breakdown
// (including the student's actual grade), and the basic submission
// information for the student's submission if it exists.
// https://canvas.instructure.com/doc/api/analytics.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.StudentID (Required) ID
//
type GetUserInACourseLevelAssignmentData struct {
	Path struct {
		CourseID  string `json:"course_id" url:"course_id,omitempty"`   //  (Required)
		StudentID string `json:"student_id" url:"student_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetUserInACourseLevelAssignmentData) GetMethod() string {
	return "GET"
}

func (t *GetUserInACourseLevelAssignmentData) GetURLPath() string {
	path := "courses/{course_id}/analytics/users/{student_id}/assignments"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{student_id}", fmt.Sprintf("%v", t.Path.StudentID))
	return path
}

func (t *GetUserInACourseLevelAssignmentData) GetQuery() (string, error) {
	return "", nil
}

func (t *GetUserInACourseLevelAssignmentData) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetUserInACourseLevelAssignmentData) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetUserInACourseLevelAssignmentData) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.StudentID == "" {
		errs = append(errs, "'Path.StudentID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetUserInACourseLevelAssignmentData) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
