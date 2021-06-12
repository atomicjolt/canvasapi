package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetCourseNickname Returns the nickname for a specific course.
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type GetCourseNickname struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`
}

func (t *GetCourseNickname) GetMethod() string {
	return "GET"
}

func (t *GetCourseNickname) GetURLPath() string {
	path := "users/self/course_nicknames/{course_id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *GetCourseNickname) GetQuery() (string, error) {
	return "", nil
}

func (t *GetCourseNickname) GetBody() (string, error) {
	return "", nil
}

func (t *GetCourseNickname) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetCourseNickname) Do(c *canvasapi.Canvas) (*models.CourseNickname, error) {
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