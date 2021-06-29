package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
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
// # Path.CourseID (Required) ID
// # Path.ModuleID (Required) ID
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.AssignmentSetID (Optional) Assignment set chosen, as specified in the mastery_paths portion of the
//    context module item response
// # Form.StudentID (Optional) Which student the selection applies to.  If not specified, current user is
//    implied.
//
type SelectMasteryPath struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ModuleID string `json:"module_id" url:"module_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Form struct {
		AssignmentSetID string `json:"assignment_set_id" url:"assignment_set_id,omitempty"` //  (Optional)
		StudentID       string `json:"student_id" url:"student_id,omitempty"`               //  (Optional)
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

func (t *SelectMasteryPath) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *SelectMasteryPath) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *SelectMasteryPath) HasErrors() error {
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

func (t *SelectMasteryPath) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
