package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// FlaggingQuestion Set a flag on a quiz question to indicate that you want to return to it
// later.
// https://canvas.instructure.com/doc/api/quiz_submission_questions.html
//
// Path Parameters:
// # QuizSubmissionID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # Attempt (Required) The attempt number of the quiz submission being taken. Note that this
//    must be the latest attempt index, as questions for earlier attempts can
//    not be modified.
// # ValidationToken (Required) The unique validation token you received when the Quiz Submission was
//    created.
// # AccessCode (Optional) Access code for the Quiz, if any.
//
type FlaggingQuestion struct {
	Path struct {
		QuizSubmissionID string `json:"quiz_submission_id"` //  (Required)
		ID               string `json:"id"`                 //  (Required)
	} `json:"path"`

	Form struct {
		Attempt         int64  `json:"attempt"`          //  (Required)
		ValidationToken string `json:"validation_token"` //  (Required)
		AccessCode      string `json:"access_code"`      //  (Optional)
	} `json:"form"`
}

func (t *FlaggingQuestion) GetMethod() string {
	return "PUT"
}

func (t *FlaggingQuestion) GetURLPath() string {
	path := "quiz_submissions/{quiz_submission_id}/questions/{id}/flag"
	path = strings.ReplaceAll(path, "{quiz_submission_id}", fmt.Sprintf("%v", t.Path.QuizSubmissionID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *FlaggingQuestion) GetQuery() (string, error) {
	return "", nil
}

func (t *FlaggingQuestion) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *FlaggingQuestion) HasErrors() error {
	errs := []string{}
	if t.Path.QuizSubmissionID == "" {
		errs = append(errs, "'QuizSubmissionID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Form.ValidationToken == "" {
		errs = append(errs, "'ValidationToken' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *FlaggingQuestion) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
