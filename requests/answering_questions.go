package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// AnsweringQuestions Provide or update an answer to one or more QuizQuestions.
// https://canvas.instructure.com/doc/api/quiz_submission_questions.html
//
// Path Parameters:
// # QuizSubmissionID (Required) ID
//
// Form Parameters:
// # Attempt (Required) The attempt number of the quiz submission being taken. Note that this
//    must be the latest attempt index, as questions for earlier attempts can
//    not be modified.
// # ValidationToken (Required) The unique validation token you received when the Quiz Submission was
//    created.
// # AccessCode (Optional) Access code for the Quiz, if any.
// # QuizQuestions (Optional) Set of question IDs and the answer value.
//
//    See {Appendix: Question Answer Formats} for the accepted answer formats
//    for each question type.
//
type AnsweringQuestions struct {
	Path struct {
		QuizSubmissionID string `json:"quiz_submission_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Attempt         int64                            `json:"attempt"`          //  (Required)
		ValidationToken string                           `json:"validation_token"` //  (Required)
		AccessCode      string                           `json:"access_code"`      //  (Optional)
		QuizQuestions   []*models.QuizSubmissionQuestion `json:"quiz_questions"`   //  (Optional)
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

func (t *AnsweringQuestions) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *AnsweringQuestions) HasErrors() error {
	errs := []string{}
	if t.Path.QuizSubmissionID == "" {
		errs = append(errs, "'QuizSubmissionID' is required")
	}
	if t.Form.ValidationToken == "" {
		errs = append(errs, "'ValidationToken' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *AnsweringQuestions) Do(c *canvasapi.Canvas) ([]*models.QuizSubmissionQuestion, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.QuizSubmissionQuestion{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
