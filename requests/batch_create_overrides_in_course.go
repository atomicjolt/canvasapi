package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// BatchCreateOverridesInCourse Creates the specified overrides for each assignment.  Handles creation in a
// transaction, so all records are created or none are.
//
// One of student_ids, group_id, or course_section_id must be present. At most
// one should be present; if multiple are present only the most specific
// (student_ids first, then group_id, then course_section_id) is used and any
// others are ignored.
//
// Errors are reported in an errors attribute, an array of errors corresponding
// to inputs.  Global errors will be reported as a single element errors array
// https://canvas.instructure.com/doc/api/assignments.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
// Form Parameters:
// # Form.AssignmentOverrides (Required) Attributes for the new assignment overrides.
//    See {api:AssignmentOverridesController#create Create an assignment override} for available
//    attributes
//
type BatchCreateOverridesInCourse struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		AssignmentOverrides []*models.AssignmentOverride `json:"assignment_overrides" url:"assignment_overrides,omitempty"` //  (Required)
	} `json:"form"`
}

func (t *BatchCreateOverridesInCourse) GetMethod() string {
	return "POST"
}

func (t *BatchCreateOverridesInCourse) GetURLPath() string {
	path := "courses/{course_id}/assignments/overrides"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *BatchCreateOverridesInCourse) GetQuery() (string, error) {
	return "", nil
}

func (t *BatchCreateOverridesInCourse) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *BatchCreateOverridesInCourse) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *BatchCreateOverridesInCourse) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Form.AssignmentOverrides == nil {
		errs = append(errs, "'Form.AssignmentOverrides' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *BatchCreateOverridesInCourse) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.AssignmentOverride, *canvasapi.PagedResource, error) {
	var err error
	var response *http.Response
	if next != nil {
		response, err = c.Send(next, t.GetMethod(), nil)
	} else {
		response, err = c.SendRequest(t)
	}

	if err != nil {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.AssignmentOverride{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, nil, err
	}

	pagedResource, err := canvasapi.ExtractPagedResource(response.Header)
	if err != nil {
		return nil, nil, err
	}

	return ret, pagedResource, nil
}
