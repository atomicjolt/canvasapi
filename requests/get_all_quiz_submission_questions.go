package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// GetAllQuizSubmissionQuestions Get a list of all the question records for this quiz submission.
//
// <b>200 OK</b> response code is returned if the request was successful.
// https://canvas.instructure.com/doc/api/quiz_submission_questions.html
//
// Path Parameters:
// # Path.QuizSubmissionID (Required) ID
//
// Query Parameters:
// # Query.Include (Optional) . Must be one of quiz_questionAssociations to include with the quiz submission question.
//
type GetAllQuizSubmissionQuestions struct {
	Path struct {
		QuizSubmissionID string `json:"quiz_submission_id" url:"quiz_submission_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of quiz_question
	} `json:"query"`
}

func (t *GetAllQuizSubmissionQuestions) GetMethod() string {
	return "GET"
}

func (t *GetAllQuizSubmissionQuestions) GetURLPath() string {
	path := "quiz_submissions/{quiz_submission_id}/questions"
	path = strings.ReplaceAll(path, "{quiz_submission_id}", fmt.Sprintf("%v", t.Path.QuizSubmissionID))
	return path
}

func (t *GetAllQuizSubmissionQuestions) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *GetAllQuizSubmissionQuestions) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetAllQuizSubmissionQuestions) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetAllQuizSubmissionQuestions) HasErrors() error {
	errs := []string{}
	if t.Path.QuizSubmissionID == "" {
		errs = append(errs, "'Path.QuizSubmissionID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"quiz_question"}, v) {
			errs = append(errs, "Include must be one of quiz_question")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetAllQuizSubmissionQuestions) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
