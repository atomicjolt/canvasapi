package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UnflaggingQuestion Remove the flag that you previously set on a quiz question after you've
// returned to it.
// https://canvas.instructure.com/doc/api/quiz_submission_questions.html
//
// Path Parameters:
// # Path.QuizSubmissionID (Required) ID
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.Attempt (Required) The attempt number of the quiz submission being taken. Note that this
//    must be the latest attempt index, as questions for earlier attempts can
//    not be modified.
// # Form.ValidationToken (Required) The unique validation token you received when the Quiz Submission was
//    created.
// # Form.AccessCode (Optional) Access code for the Quiz, if any.
//
type UnflaggingQuestion struct {
	Path struct {
		QuizSubmissionID string `json:"quiz_submission_id" url:"quiz_submission_id,omitempty"` //  (Required)
		ID               string `json:"id" url:"id,omitempty"`                                 //  (Required)
	} `json:"path"`

	Form struct {
		Attempt         int64  `json:"attempt" url:"attempt,omitempty"`                   //  (Required)
		ValidationToken string `json:"validation_token" url:"validation_token,omitempty"` //  (Required)
		AccessCode      string `json:"access_code" url:"access_code,omitempty"`           //  (Optional)
	} `json:"form"`
}

func (t *UnflaggingQuestion) GetMethod() string {
	return "PUT"
}

func (t *UnflaggingQuestion) GetURLPath() string {
	path := "quiz_submissions/{quiz_submission_id}/questions/{id}/unflag"
	path = strings.ReplaceAll(path, "{quiz_submission_id}", fmt.Sprintf("%v", t.Path.QuizSubmissionID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UnflaggingQuestion) GetQuery() (string, error) {
	return "", nil
}

func (t *UnflaggingQuestion) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UnflaggingQuestion) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UnflaggingQuestion) HasErrors() error {
	errs := []string{}
	if t.Path.QuizSubmissionID == "" {
		errs = append(errs, "'Path.QuizSubmissionID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if t.Form.ValidationToken == "" {
		errs = append(errs, "'Form.ValidationToken' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UnflaggingQuestion) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
