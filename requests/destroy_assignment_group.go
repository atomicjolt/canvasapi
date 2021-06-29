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

// DestroyAssignmentGroup Deletes the assignment group with the given id.
// https://canvas.instructure.com/doc/api/assignment_groups.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.AssignmentGroupID (Required) ID
//
// Query Parameters:
// # Query.MoveAssignmentsTo (Optional) The ID of an active Assignment Group to which the assignments that are
//    currently assigned to the destroyed Assignment Group will be assigned.
//    NOTE: If this argument is not provided, any assignments in this Assignment
//    Group will be deleted.
//
type DestroyAssignmentGroup struct {
	Path struct {
		CourseID          string `json:"course_id" url:"course_id,omitempty"`                     //  (Required)
		AssignmentGroupID string `json:"assignment_group_id" url:"assignment_group_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		MoveAssignmentsTo int64 `json:"move_assignments_to" url:"move_assignments_to,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *DestroyAssignmentGroup) GetMethod() string {
	return "DELETE"
}

func (t *DestroyAssignmentGroup) GetURLPath() string {
	path := "courses/{course_id}/assignment_groups/{assignment_group_id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_group_id}", fmt.Sprintf("%v", t.Path.AssignmentGroupID))
	return path
}

func (t *DestroyAssignmentGroup) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *DestroyAssignmentGroup) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DestroyAssignmentGroup) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DestroyAssignmentGroup) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.AssignmentGroupID == "" {
		errs = append(errs, "'Path.AssignmentGroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DestroyAssignmentGroup) Do(c *canvasapi.Canvas) (*models.AssignmentGroup, error) {
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
