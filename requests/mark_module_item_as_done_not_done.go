package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// MarkModuleItemAsDoneNotDone Mark a module item as done/not done. Use HTTP method PUT to mark as done,
// and DELETE to mark as not done.
// https://canvas.instructure.com/doc/api/modules.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.ModuleID (Required) ID
// # Path.ID (Required) ID
//
type MarkModuleItemAsDoneNotDone struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ModuleID string `json:"module_id" url:"module_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`
}

func (t *MarkModuleItemAsDoneNotDone) GetMethod() string {
	return "PUT"
}

func (t *MarkModuleItemAsDoneNotDone) GetURLPath() string {
	path := "courses/{course_id}/modules/{module_id}/items/{id}/done"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{module_id}", fmt.Sprintf("%v", t.Path.ModuleID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *MarkModuleItemAsDoneNotDone) GetQuery() (string, error) {
	return "", nil
}

func (t *MarkModuleItemAsDoneNotDone) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *MarkModuleItemAsDoneNotDone) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *MarkModuleItemAsDoneNotDone) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.ModuleID == "" {
		errs = append(errs, "'Path.ModuleID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *MarkModuleItemAsDoneNotDone) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
