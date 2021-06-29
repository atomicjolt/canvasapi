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

// GetSingleQuizQuestion Returns the quiz question with the given id
// https://canvas.instructure.com/doc/api/quiz_questions.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.QuizID (Required) ID
// # Path.ID (Required) The quiz question unique identifier.
//
type GetSingleQuizQuestion struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		QuizID   string `json:"quiz_id" url:"quiz_id,omitempty"`     //  (Required)
		ID       int64  `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`
}

func (t *GetSingleQuizQuestion) GetMethod() string {
	return "GET"
}

func (t *GetSingleQuizQuestion) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/questions/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetSingleQuizQuestion) GetQuery() (string, error) {
	return "", nil
}

func (t *GetSingleQuizQuestion) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSingleQuizQuestion) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSingleQuizQuestion) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.QuizID == "" {
		errs = append(errs, "'Path.QuizID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleQuizQuestion) Do(c *canvasapi.Canvas) (*models.QuizQuestion, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.QuizQuestion{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
