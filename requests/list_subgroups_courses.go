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

// ListSubgroupsCourses A paginated list of the immediate OutcomeGroup children of the outcome group.
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.ID (Required) ID
//
type ListSubgroupsCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`
}

func (t *ListSubgroupsCourses) GetMethod() string {
	return "GET"
}

func (t *ListSubgroupsCourses) GetURLPath() string {
	path := "courses/{course_id}/outcome_groups/{id}/subgroups"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ListSubgroupsCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ListSubgroupsCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListSubgroupsCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListSubgroupsCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListSubgroupsCourses) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.OutcomeGroup, *canvasapi.PagedResource, error) {
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
	ret := []*models.OutcomeGroup{}
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
