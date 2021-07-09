package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListCourseNicknames Returns all course nicknames you have set.
// https://canvas.instructure.com/doc/api/users.html
//
type ListCourseNicknames struct {
}

func (t *ListCourseNicknames) GetMethod() string {
	return "GET"
}

func (t *ListCourseNicknames) GetURLPath() string {
	return ""
}

func (t *ListCourseNicknames) GetQuery() (string, error) {
	return "", nil
}

func (t *ListCourseNicknames) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListCourseNicknames) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListCourseNicknames) HasErrors() error {
	return nil
}

func (t *ListCourseNicknames) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.CourseNickname, *canvasapi.PagedResource, error) {
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
	ret := []*models.CourseNickname{}
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
