package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListCollaborationsCourses A paginated list of collaborations the current user has access to in the
// context of the course provided in the url. NOTE: this only returns
// ExternalToolCollaboration type collaborations.
//
//   curl https://<canvas>/api/v1/courses/1/collaborations/
// https://canvas.instructure.com/doc/api/collaborations.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type ListCollaborationsCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListCollaborationsCourses) GetMethod() string {
	return "GET"
}

func (t *ListCollaborationsCourses) GetURLPath() string {
	path := "courses/{course_id}/collaborations"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListCollaborationsCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ListCollaborationsCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListCollaborationsCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListCollaborationsCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListCollaborationsCourses) Do(c *canvasapi.Canvas) ([]*models.Collaboration, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Collaboration{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
