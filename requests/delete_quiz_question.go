package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// DeleteQuizQuestion <b>204 No Content</b> response code is returned if the deletion was successful.
// https://canvas.instructure.com/doc/api/quiz_questions.html
//
// Path Parameters:
// # CourseID (Required) ID
// # QuizID (Required) The associated quiz's unique identifier
// # ID (Required) The quiz question's unique identifier
//
type DeleteQuizQuestion struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		QuizID   int64  `json:"quiz_id"`   //  (Required)
		ID       int64  `json:"id"`        //  (Required)
	} `json:"path"`
}

func (t *DeleteQuizQuestion) GetMethod() string {
	return "DELETE"
}

func (t *DeleteQuizQuestion) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/questions/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteQuizQuestion) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteQuizQuestion) GetBody() (string, error) {
	return "", nil
}

func (t *DeleteQuizQuestion) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteQuizQuestion) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
