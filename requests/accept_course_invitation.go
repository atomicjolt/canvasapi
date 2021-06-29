package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// AcceptCourseInvitation accepts a pending course invitation for the current user
// https://canvas.instructure.com/doc/api/enrollments.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.ID (Required) ID
//
type AcceptCourseInvitation struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`
}

func (t *AcceptCourseInvitation) GetMethod() string {
	return "POST"
}

func (t *AcceptCourseInvitation) GetURLPath() string {
	path := "courses/{course_id}/enrollments/{id}/accept"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *AcceptCourseInvitation) GetQuery() (string, error) {
	return "", nil
}

func (t *AcceptCourseInvitation) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *AcceptCourseInvitation) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *AcceptCourseInvitation) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *AcceptCourseInvitation) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
