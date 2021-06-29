package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// GetFormattedStudentNumericalAnswer Matches the intended behavior of the UI when a numerical answer is entered
// and returns the resulting formatted number
// https://canvas.instructure.com/doc/api/quiz_submission_questions.html
//
// Path Parameters:
// # Path.QuizSubmissionID (Required) ID
// # Path.ID (Required) ID
//
// Query Parameters:
// # Query.Answer (Required) no description
//
type GetFormattedStudentNumericalAnswer struct {
	Path struct {
		QuizSubmissionID string `json:"quiz_submission_id" url:"quiz_submission_id,omitempty"` //  (Required)
		ID               string `json:"id" url:"id,omitempty"`                                 //  (Required)
	} `json:"path"`

	Query struct {
		Answer float64 `json:"answer" url:"answer,omitempty"` //  (Required)
	} `json:"query"`
}

func (t *GetFormattedStudentNumericalAnswer) GetMethod() string {
	return "GET"
}

func (t *GetFormattedStudentNumericalAnswer) GetURLPath() string {
	path := "quiz_submissions/{quiz_submission_id}/questions/{id}/formatted_answer"
	path = strings.ReplaceAll(path, "{quiz_submission_id}", fmt.Sprintf("%v", t.Path.QuizSubmissionID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetFormattedStudentNumericalAnswer) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *GetFormattedStudentNumericalAnswer) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetFormattedStudentNumericalAnswer) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetFormattedStudentNumericalAnswer) HasErrors() error {
	errs := []string{}
	if t.Path.QuizSubmissionID == "" {
		errs = append(errs, "'Path.QuizSubmissionID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetFormattedStudentNumericalAnswer) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
