package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// CrossListSection Move the Section to another course.  The new course may be in a different account (department),
// but must belong to the same root account (institution).
// https://canvas.instructure.com/doc/api/sections.html
//
// Path Parameters:
// # ID (Required) ID
// # NewCourseID (Required) ID
//
type CrossListSection struct {
	Path struct {
		ID          string `json:"id"`            //  (Required)
		NewCourseID string `json:"new_course_id"` //  (Required)
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

func (t *CrossListSection) GetBody() (string, error) {
	return "", nil
}

func (t *CrossListSection) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Path.NewCourseID == "" {
		errs = append(errs, "'NewCourseID' is required")
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
