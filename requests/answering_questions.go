package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// AnsweringQuestions Provide or update an answer to one or more QuizQuestions.
// https://canvas.instructure.com/doc/api/quiz_submission_questions.html
//
// Path Parameters:
// # Path.QuizSubmissionID (Required) ID
//
// Form Parameters:
// # Form.Attempt (Required) The attempt number of the quiz submission being taken. Note that this
//    must be the latest attempt index, as questions for earlier attempts can
//    not be modified.
// # Form.ValidationToken (Required) The unique validation token you received when the Quiz Submission was
//    created.
// # Form.AccessCode (Optional) Access code for the Quiz, if any.
// # Form.QuizQuestions (Optional) Set of question IDs and the answer value.
//
//    See {Appendix: Question Answer Formats} for the accepted answer formats
//    for each question type.
//
type AnsweringQuestions struct {
	Path struct {
		QuizSubmissionID string `json:"quiz_submission_id" url:"quiz_submission_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Attempt         int64                            `json:"attempt" url:"attempt,omitempty"`                   //  (Required)
		ValidationToken string                           `json:"validation_token" url:"validation_token,omitempty"` //  (Required)
		AccessCode      string                           `json:"access_code" url:"access_code,omitempty"`           //  (Optional)
		QuizQuestions   []*models.QuizSubmissionQuestion `json:"quiz_questions" url:"quiz_questions,omitempty"`     //  (Optional)
	} `json:"form"`
}

func (t *AnsweringQuestions) GetMethod() string {
	return "POST"
}

func (t *AnsweringQuestions) GetURLPath() string {
	path := "quiz_submissions/{quiz_submission_id}/questions"
	path = strings.ReplaceAll(path, "{quiz_submission_id}", fmt.Sprintf("%v", t.Path.QuizSubmissionID))
	return path
}

func (t *AnsweringQuestions) GetQuery() (string, error) {
	return "", nil
}

func (t *AnsweringQuestions) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *AnsweringQuestions) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *AnsweringQuestions) HasErrors() error {
	errs := []string{}
	if t.Path.QuizSubmissionID == "" {
		errs = append(errs, "'Path.QuizSubmissionID' is required")
	}
	if t.Form.ValidationToken == "" {
		errs = append(errs, "'Form.ValidationToken' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *AnsweringQuestions) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.QuizSubmissionQuestion, *canvasapi.PagedResource, error) {
	var err error
	var response *http.Response
	if next != nil {
		response, err = c.Send(next, t.GetMethod(), nil)
	} else {
		response, err = c.SendRequest(t)
	}

	if err != nil {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.QuizSubmissionQuestion{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, nil, err
	}

	pagedResource, err := canvasapi.ExtractPagedResource(response.Header)
	if err != nil {
		return nil, nil, err
	}

	return ret, pagedResource, nil
}
