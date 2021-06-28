package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// RedirectToAssignmentOverrideForSection Responds with a redirect to the override for the given section, if any
// (404 otherwise).
// https://canvas.instructure.com/doc/api/assignments.html
//
// Path Parameters:
// # CourseSectionID (Required) ID
// # AssignmentID (Required) ID
//
type RedirectToAssignmentOverrideForSection struct {
	Path struct {
		CourseSectionID string `json:"course_section_id" url:"course_section_id,omitempty"` //  (Required)
		AssignmentID    string `json:"assignment_id" url:"assignment_id,omitempty"`         //  (Required)
	} `json:"path"`
}

func (t *RedirectToAssignmentOverrideForSection) GetMethod() string {
	return "GET"
}

func (t *RedirectToAssignmentOverrideForSection) GetURLPath() string {
	path := "sections/{course_section_id}/assignments/{assignment_id}/override"
	path = strings.ReplaceAll(path, "{course_section_id}", fmt.Sprintf("%v", t.Path.CourseSectionID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	return path
}

func (t *RedirectToAssignmentOverrideForSection) GetQuery() (string, error) {
	return "", nil
}

func (t *RedirectToAssignmentOverrideForSection) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *RedirectToAssignmentOverrideForSection) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *RedirectToAssignmentOverrideForSection) HasErrors() error {
	errs := []string{}
	if t.Path.CourseSectionID == "" {
		errs = append(errs, "'CourseSectionID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'AssignmentID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *RedirectToAssignmentOverrideForSection) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
