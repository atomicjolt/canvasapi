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

// ListGradingStandardsAvailableInContextCourses Returns the paginated list of grading standards for the given context that are visible to the user.
// https://canvas.instructure.com/doc/api/grading_standards.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
type ListGradingStandardsAvailableInContextCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListGradingStandardsAvailableInContextCourses) GetMethod() string {
	return "GET"
}

func (t *ListGradingStandardsAvailableInContextCourses) GetURLPath() string {
	path := "courses/{course_id}/grading_standards"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListGradingStandardsAvailableInContextCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ListGradingStandardsAvailableInContextCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListGradingStandardsAvailableInContextCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListGradingStandardsAvailableInContextCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListGradingStandardsAvailableInContextCourses) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.GradingStandard, *canvasapi.PagedResource, error) {
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
	ret := []*models.GradingStandard{}
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
