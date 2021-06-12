package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// CompleteQuizSubmissionTurnItIn Complete the quiz submission by marking it as complete and grading it. When
// the quiz submission has been marked as complete, no further modifications
// will be allowed.
//
// <b>Responses</b>
//
// * <b>200 OK</b> if the request was successful
// * <b>403 Forbidden</b> if an invalid access code is specified
// * <b>403 Forbidden</b> if the Quiz's IP filter restriction does not pass
// * <b>403 Forbidden</b> if an invalid token is specified
// * <b>400 Bad Request</b> if the QS is already complete
// * <b>400 Bad Request</b> if the attempt parameter is missing
// * <b>400 Bad Request</b> if the attempt parameter is not the latest attempt
// https://canvas.instructure.com/doc/api/quiz_submissions.html
//
// Path Parameters:
// # CourseID (Required) ID
// # QuizID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # Attempt (Required) The attempt number of the quiz submission that should be completed. Note
//    that this must be the latest attempt index, as earlier attempts can not
//    be modified.
// # ValidationToken (Required) The unique validation token you received when this Quiz Submission was
//    created.
// # AccessCode (Optional) Access code for the Quiz, if any.
//
type CompleteQuizSubmissionTurnItIn struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		QuizID   string `json:"quiz_id"`   //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`

	Form struct {
		Attempt         int64  `json:"attempt"`          //  (Required)
		ValidationToken string `json:"validation_token"` //  (Required)
		AccessCode      string `json:"access_code"`      //  (Optional)
	} `json:"form"`
}

func (t *CompleteQuizSubmissionTurnItIn) GetMethod() string {
	return "POST"
}

func (t *CompleteQuizSubmissionTurnItIn) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/submissions/{id}/complete"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *CompleteQuizSubmissionTurnItIn) GetQuery() (string, error) {
	return "", nil
}

func (t *CompleteQuizSubmissionTurnItIn) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CompleteQuizSubmissionTurnItIn) HasErrors() error {
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
	if t.Form.ValidationToken == "" {
		errs = append(errs, "'ValidationToken' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CompleteQuizSubmissionTurnItIn) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
