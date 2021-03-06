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

// ListConferencesCourses Retrieve the paginated list of conferences for this context
//
// This API returns a JSON object containing the list of conferences,
// the key for the list of conferences is "conferences"
// https://canvas.instructure.com/doc/api/conferences.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
type ListConferencesCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListConferencesCourses) GetMethod() string {
	return "GET"
}

func (t *ListConferencesCourses) GetURLPath() string {
	path := "courses/{course_id}/conferences"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListConferencesCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ListConferencesCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListConferencesCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListConferencesCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListConferencesCourses) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.Conference, *canvasapi.PagedResource, error) {
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
	ret := []*models.Conference{}
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
