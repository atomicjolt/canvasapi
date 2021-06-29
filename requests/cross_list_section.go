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

// CrossListSection Move the Section to another course.  The new course may be in a different account (department),
// but must belong to the same root account (institution).
// https://canvas.instructure.com/doc/api/sections.html
//
// Path Parameters:
// # Path.ID (Required) ID
// # Path.NewCourseID (Required) ID
//
type CrossListSection struct {
	Path struct {
		ID          string `json:"id" url:"id,omitempty"`                       //  (Required)
		NewCourseID string `json:"new_course_id" url:"new_course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *CrossListSection) GetMethod() string {
	return "POST"
}

func (t *CrossListSection) GetURLPath() string {
	path := "sections/{id}/crosslist/{new_course_id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	path = strings.ReplaceAll(path, "{new_course_id}", fmt.Sprintf("%v", t.Path.NewCourseID))
	return path
}

func (t *CrossListSection) GetQuery() (string, error) {
	return "", nil
}

func (t *CrossListSection) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *CrossListSection) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *CrossListSection) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if t.Path.NewCourseID == "" {
		errs = append(errs, "'Path.NewCourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CrossListSection) Do(c *canvasapi.Canvas) (*models.Section, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Section{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
