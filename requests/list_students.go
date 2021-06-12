package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListStudents Returns the paginated list of students enrolled in this course.
//
// DEPRECATED: Please use the {api:CoursesController#users course users} endpoint
// and pass "student" as the enrollment_type.
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type ListStudents struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`
}

func (t *ListStudents) GetMethod() string {
	return "GET"
}

func (t *ListStudents) GetURLPath() string {
	path := "courses/{course_id}/students"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListStudents) GetQuery() (string, error) {
	return "", nil
}

func (t *ListStudents) GetBody() (string, error) {
	return "", nil
}

func (t *ListStudents) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListStudents) Do(c *canvasapi.Canvas) ([]*models.User, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.User{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
