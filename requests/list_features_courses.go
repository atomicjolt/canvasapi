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

// ListFeaturesCourses A paginated list of all features that apply to a given Account, Course, or User.
// https://canvas.instructure.com/doc/api/feature_flags.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
type ListFeaturesCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListFeaturesCourses) GetMethod() string {
	return "GET"
}

func (t *ListFeaturesCourses) GetURLPath() string {
	path := "courses/{course_id}/features"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListFeaturesCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ListFeaturesCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListFeaturesCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListFeaturesCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListFeaturesCourses) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.Feature, *canvasapi.PagedResource, error) {
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
	ret := []*models.Feature{}
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
