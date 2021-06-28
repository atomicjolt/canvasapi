package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetCurrentQuizSubmissionTimes Get the current timing data for the quiz attempt, both the end_at timestamp
// and the time_left parameter.
//
// <b>Responses</b>
//
// * <b>200 OK</b> if the request was successful
// https://canvas.instructure.com/doc/api/quiz_submissions.html
//
// Path Parameters:
// # CourseID (Required) ID
// # QuizID (Required) ID
// # ID (Required) ID
//
type GetCurrentQuizSubmissionTimes struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		QuizID   string `json:"quiz_id" url:"quiz_id,omitempty"`     //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`
}

func (t *GetCurrentQuizSubmissionTimes) GetMethod() string {
	return "GET"
}

func (t *GetCurrentQuizSubmissionTimes) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/submissions/{id}/time"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetCurrentQuizSubmissionTimes) GetQuery() (string, error) {
	return "", nil
}

func (t *GetCurrentQuizSubmissionTimes) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetCurrentQuizSubmissionTimes) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetCurrentQuizSubmissionTimes) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.QuizID == "" {
		errs = append(errs, "'QuizID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetCurrentQuizSubmissionTimes) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
