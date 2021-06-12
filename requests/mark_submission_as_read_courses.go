package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// MarkSubmissionAsReadCourses No request fields are necessary.
//
// On success, the response will be 204 No Content with an empty body.
// https://canvas.instructure.com/doc/api/submissions.html
//
// Path Parameters:
// # CourseID (Required) ID
// # AssignmentID (Required) ID
// # UserID (Required) ID
//
type MarkSubmissionAsReadCourses struct {
	Path struct {
		CourseID     string `json:"course_id"`     //  (Required)
		AssignmentID string `json:"assignment_id"` //  (Required)
		UserID       string `json:"user_id"`       //  (Required)
	} `json:"path"`
}

func (t *MarkSubmissionAsReadCourses) GetMethod() string {
	return "PUT"
}

func (t *MarkSubmissionAsReadCourses) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/submissions/{user_id}/read"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *MarkSubmissionAsReadCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *MarkSubmissionAsReadCourses) GetBody() (string, error) {
	return "", nil
}

func (t *MarkSubmissionAsReadCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
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

func (t *MarkSubmissionAsReadCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}