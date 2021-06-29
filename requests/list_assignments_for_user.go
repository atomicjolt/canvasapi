package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ListAssignmentsForUser Returns the paginated list of assignments for the specified user if the current user has rights to view.
// See {api:AssignmentsApiController#index List assignments} for valid arguments.
// https://canvas.instructure.com/doc/api/assignments.html
//
// Path Parameters:
// # Path.UserID (Required) ID
// # Path.CourseID (Required) ID
//
type ListAssignmentsForUser struct {
	Path struct {
		UserID   string `json:"user_id" url:"user_id,omitempty"`     //  (Required)
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListAssignmentsForUser) GetMethod() string {
	return "GET"
}

func (t *ListAssignmentsForUser) GetURLPath() string {
	path := "users/{user_id}/courses/{course_id}/assignments"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListAssignmentsForUser) GetQuery() (string, error) {
	return "", nil
}

func (t *ListAssignmentsForUser) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListAssignmentsForUser) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListAssignmentsForUser) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListAssignmentsForUser) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
