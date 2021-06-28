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

// AddsLastAttendedDateToStudentEnrollmentInCourse
// https://canvas.instructure.com/doc/api/enrollments.html
//
// Path Parameters:
// # CourseID (Required) ID
// # UserID (Required) ID
//
type AddsLastAttendedDateToStudentEnrollmentInCourse struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		UserID   string `json:"user_id" url:"user_id,omitempty"`     //  (Required)
	} `json:"path"`
}

func (t *AddsLastAttendedDateToStudentEnrollmentInCourse) GetMethod() string {
	return "PUT"
}

func (t *AddsLastAttendedDateToStudentEnrollmentInCourse) GetURLPath() string {
	path := "courses/{course_id}/users/{user_id}/last_attended"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *AddsLastAttendedDateToStudentEnrollmentInCourse) GetQuery() (string, error) {
	return "", nil
}

func (t *AddsLastAttendedDateToStudentEnrollmentInCourse) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *AddsLastAttendedDateToStudentEnrollmentInCourse) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *AddsLastAttendedDateToStudentEnrollmentInCourse) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *AddsLastAttendedDateToStudentEnrollmentInCourse) Do(c *canvasapi.Canvas) (*models.Enrollment, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Enrollment{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
