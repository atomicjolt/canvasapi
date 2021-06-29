package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// DeleteQuestionGroup Delete a question group
//
// <b>204 No Content<b> response code is returned if the deletion was successful.
// https://canvas.instructure.com/doc/api/quiz_question_groups.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.QuizID (Required) ID
// # Path.ID (Required) ID
//
type DeleteQuestionGroup struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		QuizID   string `json:"quiz_id" url:"quiz_id,omitempty"`     //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`
}

func (t *DeleteQuestionGroup) GetMethod() string {
	return "DELETE"
}

func (t *DeleteQuestionGroup) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/groups/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteQuestionGroup) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteQuestionGroup) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteQuestionGroup) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteQuestionGroup) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.QuizID == "" {
		errs = append(errs, "'Path.QuizID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteQuestionGroup) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
