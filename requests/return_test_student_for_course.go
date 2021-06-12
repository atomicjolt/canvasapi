package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ReturnTestStudentForCourse Returns information for a test student in this course. Creates a test
// student if one does not already exist for the course. The caller must have
// permission to access the course's student view.
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type ReturnTestStudentForCourse struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`
}

func (t *ReturnTestStudentForCourse) GetMethod() string {
	return "GET"
}

func (t *ReturnTestStudentForCourse) GetURLPath() string {
	path := "courses/{course_id}/student_view_student"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ReturnTestStudentForCourse) GetQuery() (string, error) {
	return "", nil
}

func (t *ReturnTestStudentForCourse) GetBody() (string, error) {
	return "", nil
}

func (t *ReturnTestStudentForCourse) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ReturnTestStudentForCourse) Do(c *canvasapi.Canvas) (*models.User, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.User{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
