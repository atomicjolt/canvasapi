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

// BatchRetrieveOverridesInCourse Returns a list of specified overrides in this course, providing
// they target sections/groups/students visible to the current user.
// Returns null elements in the list for requests that were not found.
// https://canvas.instructure.com/doc/api/assignments.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # AssignmentOverrides (Required) Ids of overrides to retrieve
// # AssignmentOverrides (Required) Ids of assignments for each override
//
type BatchRetrieveOverridesInCourse struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		AssignmentOverrides struct {
			ID           []string `json:"id" url:"id,omitempty"`                       //  (Required)
			AssignmentID []string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
		} `json:"assignment_overrides" url:"assignment_overrides,omitempty"`
	} `json:"query"`
}

func (t *BatchRetrieveOverridesInCourse) GetMethod() string {
	return "GET"
}

func (t *BatchRetrieveOverridesInCourse) GetURLPath() string {
	path := "courses/{course_id}/assignments/overrides"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *BatchRetrieveOverridesInCourse) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *BatchRetrieveOverridesInCourse) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *BatchRetrieveOverridesInCourse) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *BatchRetrieveOverridesInCourse) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Query.AssignmentOverrides.ID == nil {
		errs = append(errs, "'AssignmentOverrides' is required")
	}
	if t.Query.AssignmentOverrides.AssignmentID == nil {
		errs = append(errs, "'AssignmentOverrides' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *BatchRetrieveOverridesInCourse) Do(c *canvasapi.Canvas) ([]*models.AssignmentOverride, error) {
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
