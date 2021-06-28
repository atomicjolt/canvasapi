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

// RemoveCourseNickname Remove the nickname for the given course.
// Subsequent course API calls will return the actual name for the course.
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type RemoveCourseNickname struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *RemoveCourseNickname) GetMethod() string {
	return "DELETE"
}

func (t *RemoveCourseNickname) GetURLPath() string {
	path := "users/self/course_nicknames/{course_id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *RemoveCourseNickname) GetQuery() (string, error) {
	return "", nil
}

func (t *RemoveCourseNickname) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *RemoveCourseNickname) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *RemoveCourseNickname) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *RemoveCourseNickname) Do(c *canvasapi.Canvas) (*models.CourseNickname, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.CourseNickname{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
