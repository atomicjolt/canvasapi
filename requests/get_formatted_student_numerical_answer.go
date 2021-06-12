package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// GetFormattedStudentNumericalAnswer Matches the intended behavior of the UI when a numerical answer is entered
// and returns the resulting formatted number
// https://canvas.instructure.com/doc/api/quiz_submission_questions.html
//
// Path Parameters:
// # QuizSubmissionID (Required) ID
// # ID (Required) ID
//
// Query Parameters:
// # Answer (Required) no description
//
type GetFormattedStudentNumericalAnswer struct {
	Path struct {
		QuizSubmissionID string `json:"quiz_submission_id"` //  (Required)
		ID               string `json:"id"`                 //  (Required)
	} `json:"path"`

	Query struct {
		Answer string `json:"answer"` //  (Required)
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
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetFormattedStudentNumericalAnswer) GetBody() (string, error) {
	return "", nil
}

func (t *GetFormattedStudentNumericalAnswer) HasErrors() error {
	errs := []string{}
	if t.Path.QuizSubmissionID == "" {
		errs = append(errs, "'QuizSubmissionID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Query.Answer == "" {
		errs = append(errs, "'Answer' is required")
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
