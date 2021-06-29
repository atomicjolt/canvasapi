package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// MarkModuleItemRead Fulfills "must view" requirement for a module item. It is generally not necessary to do this explicitly,
// but it is provided for applications that need to access external content directly (bypassing the html_url
// redirect that normally allows Canvas to fulfill "must view" requirements).
//
// This endpoint cannot be used to complete requirements on locked or unpublished module items.
// https://canvas.instructure.com/doc/api/modules.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.ModuleID (Required) ID
// # Path.ID (Required) ID
//
type MarkModuleItemRead struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ModuleID string `json:"module_id" url:"module_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`
}

func (t *MarkModuleItemRead) GetMethod() string {
	return "POST"
}

func (t *MarkModuleItemRead) GetURLPath() string {
	path := "courses/{course_id}/modules/{module_id}/items/{id}/mark_read"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{module_id}", fmt.Sprintf("%v", t.Path.ModuleID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *MarkModuleItemRead) GetQuery() (string, error) {
	return "", nil
}

func (t *MarkModuleItemRead) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *MarkModuleItemRead) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *MarkModuleItemRead) HasErrors() error {
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

func (t *MarkModuleItemRead) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
