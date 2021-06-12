package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// RedirectToAssignmentOverrideForGroup Responds with a redirect to the override for the given group, if any
// (404 otherwise).
// https://canvas.instructure.com/doc/api/assignments.html
//
// Path Parameters:
// # GroupID (Required) ID
// # AssignmentID (Required) ID
//
type RedirectToAssignmentOverrideForGroup struct {
	Path struct {
		GroupID      string `json:"group_id"`      //  (Required)
		AssignmentID string `json:"assignment_id"` //  (Required)
	} `json:"path"`
}

func (t *RedirectToAssignmentOverrideForGroup) GetMethod() string {
	return "GET"
}

func (t *RedirectToAssignmentOverrideForGroup) GetURLPath() string {
	path := "groups/{group_id}/assignments/{assignment_id}/override"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	return path
}

func (t *RedirectToAssignmentOverrideForGroup) GetQuery() (string, error) {
	return "", nil
}

func (t *RedirectToAssignmentOverrideForGroup) GetBody() (string, error) {
	return "", nil
}

func (t *RedirectToAssignmentOverrideForGroup) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'AssignmentID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *RedirectToAssignmentOverrideForGroup) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
