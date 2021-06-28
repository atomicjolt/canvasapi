package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// BatchUpdateOverridesInCourse Updates a list of specified overrides for each assignment.  Handles overrides
// in a transaction, so either all updates are applied or none.
// See {api:AssignmentOverridesController#update Update an assignment override} for
// available attributes.
//
// All current overridden values must be supplied if they are to be retained;
// e.g. if due_at was overridden, but this PUT omits a value for due_at,
// due_at will no longer be overridden. If the override is adhoc and
// student_ids is not supplied, the target override set is unchanged. Target
// override sets cannot be changed for group or section overrides.
//
// Errors are reported in an errors attribute, an array of errors corresponding
// to inputs.  Global errors will be reported as a single element errors array
// https://canvas.instructure.com/doc/api/assignments.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # AssignmentOverrides (Required) Attributes for the updated overrides.
//
type BatchUpdateOverridesInCourse struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		AssignmentOverrides []*models.AssignmentOverride `json:"assignment_overrides" url:"assignment_overrides,omitempty"` //  (Required)
	} `json:"form"`
}

func (t *BatchUpdateOverridesInCourse) GetMethod() string {
	return "PUT"
}

func (t *BatchUpdateOverridesInCourse) GetURLPath() string {
	path := "courses/{course_id}/assignments/overrides"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *BatchUpdateOverridesInCourse) GetQuery() (string, error) {
	return "", nil
}

func (t *BatchUpdateOverridesInCourse) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *BatchUpdateOverridesInCourse) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *BatchUpdateOverridesInCourse) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Form.AssignmentOverrides == nil {
		errs = append(errs, "'AssignmentOverrides' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *BatchUpdateOverridesInCourse) Do(c *canvasapi.Canvas) ([]*models.AssignmentOverride, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.AssignmentOverride{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
