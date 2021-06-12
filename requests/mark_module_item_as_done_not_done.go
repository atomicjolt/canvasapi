package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// MarkModuleItemAsDoneNotDone Mark a module item as done/not done. Use HTTP method PUT to mark as done,
// and DELETE to mark as not done.
// https://canvas.instructure.com/doc/api/modules.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ModuleID (Required) ID
// # ID (Required) ID
//
type MarkModuleItemAsDoneNotDone struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ModuleID string `json:"module_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
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

func (t *MarkModuleItemAsDoneNotDone) GetBody() (string, error) {
	return "", nil
}

func (t *MarkModuleItemAsDoneNotDone) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ModuleID == "" {
		errs = append(errs, "'ModuleID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
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
