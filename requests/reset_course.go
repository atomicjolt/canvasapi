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

// ResetCourse Deletes the current course, and creates a new equivalent course with
// no content, but all sections and users moved over.
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type ResetCourse struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ResetCourse) GetMethod() string {
	return "POST"
}

func (t *ResetCourse) GetURLPath() string {
	path := "courses/{course_id}/reset_content"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ResetCourse) GetQuery() (string, error) {
	return "", nil
}

func (t *ResetCourse) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ResetCourse) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ResetCourse) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ResetCourse) Do(c *canvasapi.Canvas) (*models.Course, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Course{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
