package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// RejectCourseInvitation rejects a pending course invitation for the current user
// https://canvas.instructure.com/doc/api/enrollments.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
type RejectCourseInvitation struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`
}

func (t *RejectCourseInvitation) GetMethod() string {
	return "POST"
}

func (t *RejectCourseInvitation) GetURLPath() string {
	path := "courses/{course_id}/enrollments/{id}/reject"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *RejectCourseInvitation) GetQuery() (string, error) {
	return "", nil
}

func (t *RejectCourseInvitation) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *RejectCourseInvitation) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *RejectCourseInvitation) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *RejectCourseInvitation) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
