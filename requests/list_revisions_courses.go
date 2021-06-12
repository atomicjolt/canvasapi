package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListRevisionsCourses A paginated list of the revisions of a page. Callers must have update rights on the page in order to see page history.
// https://canvas.instructure.com/doc/api/pages.html
//
// Path Parameters:
// # CourseID (Required) ID
// # Url (Required) ID
//
type ListRevisionsCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		Url      string `json:"url"`       //  (Required)
	} `json:"path"`
}

func (t *ListRevisionsCourses) GetMethod() string {
	return "GET"
}

func (t *ListRevisionsCourses) GetURLPath() string {
	path := "courses/{course_id}/pages/{url}/revisions"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{url}", fmt.Sprintf("%v", t.Path.Url))
	return path
}

func (t *ListRevisionsCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ListRevisionsCourses) GetBody() (string, error) {
	return "", nil
}

func (t *ListRevisionsCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.Url == "" {
		errs = append(errs, "'Url' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListRevisionsCourses) Do(c *canvasapi.Canvas) ([]*models.PageRevision, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.PageRevision{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
