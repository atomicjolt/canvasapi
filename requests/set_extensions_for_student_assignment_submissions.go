package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// SetExtensionsForStudentAssignmentSubmissions <b>Responses</b>
//
// * <b>200 OK</b> if the request was successful
// * <b>403 Forbidden</b> if you are not allowed to extend assignments for this course
// * <b>400 Bad Request</b> if any of the extensions are invalid
// https://canvas.instructure.com/doc/api/assignment_extensions.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.AssignmentID (Required) ID
//
// Form Parameters:
// # Form.AssignmentExtensions.UserID (Required) The ID of the user we want to add assignment extensions for.
// # Form.AssignmentExtensions.ExtraAttempts (Required) Number of times the student is allowed to re-take the assignment over the
//    limit.
//
type SetExtensionsForStudentAssignmentSubmissions struct {
	Path struct {
		CourseID     string `json:"course_id" url:"course_id,omitempty"`         //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		AssignmentExtensions struct {
			UserID        []string `json:"user_id" url:"user_id,omitempty"`               //  (Required)
			ExtraAttempts []string `json:"extra_attempts" url:"extra_attempts,omitempty"` //  (Required)
		} `json:"assignment_extensions" url:"assignment_extensions,omitempty"`
	} `json:"form"`
}

func (t *SetExtensionsForStudentAssignmentSubmissions) GetMethod() string {
	return "POST"
}

func (t *SetExtensionsForStudentAssignmentSubmissions) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/extensions"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	return path
}

func (t *SetExtensionsForStudentAssignmentSubmissions) GetQuery() (string, error) {
	return "", nil
}

func (t *SetExtensionsForStudentAssignmentSubmissions) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *SetExtensionsForStudentAssignmentSubmissions) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *SetExtensionsForStudentAssignmentSubmissions) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'Path.AssignmentID' is required")
	}
	if t.Form.AssignmentExtensions.UserID == nil {
		errs = append(errs, "'Form.AssignmentExtensions.UserID' is required")
	}
	if t.Form.AssignmentExtensions.ExtraAttempts == nil {
		errs = append(errs, "'Form.AssignmentExtensions.ExtraAttempts' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *SetExtensionsForStudentAssignmentSubmissions) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
