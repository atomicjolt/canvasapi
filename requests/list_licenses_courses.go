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

// ListLicensesCourses A paginated list of licenses that can be applied
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
type ListLicensesCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListLicensesCourses) GetMethod() string {
	return "GET"
}

func (t *ListLicensesCourses) GetURLPath() string {
	path := "courses/{course_id}/content_licenses"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListLicensesCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ListLicensesCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListLicensesCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListLicensesCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListLicensesCourses) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.License, *canvasapi.PagedResource, error) {
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
	ret := []*models.License{}
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
