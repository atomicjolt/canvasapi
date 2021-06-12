package requests

import (
	"fmt"
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
// # CourseID (Required) ID
// # StudentID (Required) ID
//
type GetUserInACourseLevelAssignmentData struct {
	Path struct {
		CourseID  string `json:"course_id"`  //  (Required)
		StudentID string `json:"student_id"` //  (Required)
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

func (t *GetUserInACourseLevelAssignmentData) GetBody() (string, error) {
	return "", nil
}

func (t *GetUserInACourseLevelAssignmentData) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.StudentID == "" {
		errs = append(errs, "'StudentID' is required")
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
