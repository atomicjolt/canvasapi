package requests

import (
	"encoding/json"
	"io/ioutil"

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

func (t *ListCourseNicknames) GetBody() (string, error) {
	return "", nil
}

func (t *ListCourseNicknames) HasErrors() error {
	return nil
}

func (t *ListCourseNicknames) Do(c *canvasapi.Canvas) ([]*models.CourseNickname, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.CourseNickname{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
