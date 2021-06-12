package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// EditAssignmentGroup Modify an existing Assignment Group.
// Accepts the same parameters as Assignment Group creation
// https://canvas.instructure.com/doc/api/assignment_groups.html
//
// Path Parameters:
// # CourseID (Required) ID
// # AssignmentGroupID (Required) ID
//
type EditAssignmentGroup struct {
	Path struct {
		CourseID          string `json:"course_id"`           //  (Required)
		AssignmentGroupID string `json:"assignment_group_id"` //  (Required)
	} `json:"path"`
}

func (t *EditAssignmentGroup) GetMethod() string {
	return "PUT"
}

func (t *EditAssignmentGroup) GetURLPath() string {
	path := "courses/{course_id}/assignment_groups/{assignment_group_id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_group_id}", fmt.Sprintf("%v", t.Path.AssignmentGroupID))
	return path
}

func (t *EditAssignmentGroup) GetQuery() (string, error) {
	return "", nil
}

func (t *EditAssignmentGroup) GetBody() (string, error) {
	return "", nil
}

func (t *EditAssignmentGroup) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.AssignmentGroupID == "" {
		errs = append(errs, "'AssignmentGroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *EditAssignmentGroup) Do(c *canvasapi.Canvas) (*models.AssignmentGroup, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.AssignmentGroup{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
