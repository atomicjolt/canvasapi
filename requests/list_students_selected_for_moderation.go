package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListStudentsSelectedForModeration Returns a paginated list of students selected for moderation
// https://canvas.instructure.com/doc/api/moderated_grading.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.AssignmentID (Required) ID
//
type ListStudentsSelectedForModeration struct {
	Path struct {
		CourseID     string `json:"course_id" url:"course_id,omitempty"`         //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListStudentsSelectedForModeration) GetMethod() string {
	return "GET"
}

func (t *ListStudentsSelectedForModeration) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/moderated_students"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	return path
}

func (t *ListStudentsSelectedForModeration) GetQuery() (string, error) {
	return "", nil
}

func (t *ListStudentsSelectedForModeration) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListStudentsSelectedForModeration) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListStudentsSelectedForModeration) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'Path.AssignmentID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListStudentsSelectedForModeration) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.User, *canvasapi.PagedResource, error) {
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
	ret := []*models.User{}
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
