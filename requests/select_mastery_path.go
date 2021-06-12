package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// SelectMasteryPath Select a mastery path when module item includes several possible paths.
// Requires Mastery Paths feature to be enabled.  Returns a compound document
// with the assignments included in the given path and any module items
// related to those assignments
// https://canvas.instructure.com/doc/api/modules.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ModuleID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # AssignmentSetID (Optional) Assignment set chosen, as specified in the mastery_paths portion of the
//    context module item response
// # StudentID (Optional) Which student the selection applies to.  If not specified, current user is
//    implied.
//
type SelectMasteryPath struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ModuleID string `json:"module_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`

	Form struct {
		AssignmentSetID string `json:"assignment_set_id"` //  (Optional)
		StudentID       string `json:"student_id"`        //  (Optional)
	} `json:"form"`
}

func (t *SelectMasteryPath) GetMethod() string {
	return "POST"
}

func (t *SelectMasteryPath) GetURLPath() string {
	path := "courses/{course_id}/modules/{module_id}/items/{id}/select_mastery_path"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{module_id}", fmt.Sprintf("%v", t.Path.ModuleID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *SelectMasteryPath) GetQuery() (string, error) {
	return "", nil
}

func (t *SelectMasteryPath) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *SelectMasteryPath) HasErrors() error {
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

func (t *SelectMasteryPath) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
