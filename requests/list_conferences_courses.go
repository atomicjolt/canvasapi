package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
// # CourseID (Required) ID
//
type ListConferencesCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
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

func (t *ListConferencesCourses) GetBody() (string, error) {
	return "", nil
}

func (t *ListConferencesCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListConferencesCourses) Do(c *canvasapi.Canvas) ([]*models.Conference, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Conference{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
