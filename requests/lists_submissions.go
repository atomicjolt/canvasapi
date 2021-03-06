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

// ListsSubmissions Gives a nested list of submission versions
// https://canvas.instructure.com/doc/api/gradebook_history.html
//
// Path Parameters:
// # Path.CourseID (Required) The id of the contextual course for this API call
// # Path.Date (Required) The date for which you would like to see submissions
// # Path.GraderID (Required) The ID of the grader for which you want to see submissions
// # Path.AssignmentID (Required) The ID of the assignment for which you want to see submissions
//
type ListsSubmissions struct {
	Path struct {
		CourseID     int64  `json:"course_id" url:"course_id,omitempty"`         //  (Required)
		Date         string `json:"date" url:"date,omitempty"`                   //  (Required)
		GraderID     int64  `json:"grader_id" url:"grader_id,omitempty"`         //  (Required)
		AssignmentID int64  `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListsSubmissions) GetMethod() string {
	return "GET"
}

func (t *ListsSubmissions) GetURLPath() string {
	path := "courses/{course_id}/gradebook_history/{date}/graders/{grader_id}/assignments/{assignment_id}/submissions"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{date}", fmt.Sprintf("%v", t.Path.Date))
	path = strings.ReplaceAll(path, "{grader_id}", fmt.Sprintf("%v", t.Path.GraderID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	return path
}

func (t *ListsSubmissions) GetQuery() (string, error) {
	return "", nil
}

func (t *ListsSubmissions) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListsSubmissions) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListsSubmissions) HasErrors() error {
	errs := []string{}
	if t.Path.Date == "" {
		errs = append(errs, "'Path.Date' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListsSubmissions) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.SubmissionHistory, *canvasapi.PagedResource, error) {
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
	ret := []*models.SubmissionHistory{}
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
