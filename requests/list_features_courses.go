package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListFeaturesCourses A paginated list of all features that apply to a given Account, Course, or User.
// https://canvas.instructure.com/doc/api/feature_flags.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type ListFeaturesCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
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

func (t *ListFeaturesCourses) GetBody() (string, error) {
	return "", nil
}

func (t *ListFeaturesCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListFeaturesCourses) Do(c *canvasapi.Canvas) ([]*models.Feature, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Feature{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
