package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetUserProgress Return progress information for the user and course
//
// You can supply +self+ as the user_id to query your own progress in a course. To query another user's progress,
// you must be a teacher in the course, an administrator, or a linked observer of the user.
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # CourseID (Required) ID
// # UserID (Required) ID
//
type GetUserProgress struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		UserID   string `json:"user_id"`   //  (Required)
	} `json:"path"`
}

func (t *GetUserProgress) GetMethod() string {
	return "GET"
}

func (t *GetUserProgress) GetURLPath() string {
	path := "courses/{course_id}/users/{user_id}/progress"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *GetUserProgress) GetQuery() (string, error) {
	return "", nil
}

func (t *GetUserProgress) GetBody() (string, error) {
	return "", nil
}

func (t *GetUserProgress) HasErrors() error {
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

func (t *GetUserProgress) Do(c *canvasapi.Canvas) (*models.CourseProgress, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.CourseProgress{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
