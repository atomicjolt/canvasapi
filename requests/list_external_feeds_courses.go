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

// ListExternalFeedsCourses Returns the paginated list of External Feeds this course or group.
// https://canvas.instructure.com/doc/api/announcement_external_feeds.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
type ListExternalFeedsCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListExternalFeedsCourses) GetMethod() string {
	return "GET"
}

func (t *ListExternalFeedsCourses) GetURLPath() string {
	path := "courses/{course_id}/external_feeds"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListExternalFeedsCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ListExternalFeedsCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListExternalFeedsCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListExternalFeedsCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListExternalFeedsCourses) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.ExternalFeed, *canvasapi.PagedResource, error) {
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
	ret := []*models.ExternalFeed{}
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
