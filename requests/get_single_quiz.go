package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetSingleQuiz Returns the quiz with the given id.
// https://canvas.instructure.com/doc/api/quizzes.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
type GetSingleQuiz struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`
}

func (t *GetSingleQuiz) GetMethod() string {
	return "GET"
}

func (t *GetSingleQuiz) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetSingleQuiz) GetQuery() (string, error) {
	return "", nil
}

func (t *GetSingleQuiz) GetBody() (string, error) {
	return "", nil
}

func (t *GetSingleQuiz) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleQuiz) Do(c *canvasapi.Canvas) (*models.Quiz, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Quiz{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}