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

// ListRevisionsCourses A paginated list of the revisions of a page. Callers must have update rights on the page in order to see page history.
// https://canvas.instructure.com/doc/api/pages.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.Url (Required) ID
//
type ListRevisionsCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		Url      string `json:"url" url:"url,omitempty"`             //  (Required)
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

func (t *ListRevisionsCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListRevisionsCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListRevisionsCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.Url == "" {
		errs = append(errs, "'Path.Url' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListRevisionsCourses) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.PageRevision, *canvasapi.PagedResource, error) {
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
	ret := []*models.PageRevision{}
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
