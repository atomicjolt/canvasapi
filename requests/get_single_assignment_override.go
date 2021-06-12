package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetSingleAssignmentOverride Returns details of the the override with the given id.
// https://canvas.instructure.com/doc/api/assignments.html
//
// Path Parameters:
// # CourseID (Required) ID
// # AssignmentID (Required) ID
// # ID (Required) ID
//
type GetSingleAssignmentOverride struct {
	Path struct {
		CourseID     string `json:"course_id"`     //  (Required)
		AssignmentID string `json:"assignment_id"` //  (Required)
		ID           string `json:"id"`            //  (Required)
	} `json:"path"`
}

func (t *GetSingleAssignmentOverride) GetMethod() string {
	return "GET"
}

func (t *GetSingleAssignmentOverride) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/overrides/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetSingleAssignmentOverride) GetQuery() (string, error) {
	return "", nil
}

func (t *GetSingleAssignmentOverride) GetBody() (string, error) {
	return "", nil
}

func (t *GetSingleAssignmentOverride) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'AssignmentID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleAssignmentOverride) Do(c *canvasapi.Canvas) (*models.AssignmentOverride, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.AssignmentOverride{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
