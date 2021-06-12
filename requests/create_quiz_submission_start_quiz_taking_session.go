package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// CreateQuizSubmissionStartQuizTakingSession Start taking a Quiz by creating a QuizSubmission which you can use to answer
// questions and submit your answers.
//
// <b>Responses</b>
//
// * <b>200 OK</b> if the request was successful
// * <b>400 Bad Request</b> if the quiz is locked
// * <b>403 Forbidden</b> if an invalid access code is specified
// * <b>403 Forbidden</b> if the Quiz's IP filter restriction does not pass
// * <b>409 Conflict</b> if a QuizSubmission already exists for this user and quiz
// https://canvas.instructure.com/doc/api/quiz_submissions.html
//
// Path Parameters:
// # CourseID (Required) ID
// # QuizID (Required) ID
//
// Form Parameters:
// # AccessCode (Optional) Access code for the Quiz, if any.
// # Preview (Optional) Whether this should be a preview QuizSubmission and not count towards
//    the user's course record. Teachers only.
//
type CreateQuizSubmissionStartQuizTakingSession struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		QuizID   string `json:"quiz_id"`   //  (Required)
	} `json:"path"`

	Form struct {
		AccessCode string `json:"access_code"` //  (Optional)
		Preview    bool   `json:"preview"`     //  (Optional)
	} `json:"form"`
}

func (t *CreateQuizSubmissionStartQuizTakingSession) GetMethod() string {
	return "POST"
}

func (t *CreateQuizSubmissionStartQuizTakingSession) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/submissions"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	return path
}

func (t *CreateQuizSubmissionStartQuizTakingSession) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateQuizSubmissionStartQuizTakingSession) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateQuizSubmissionStartQuizTakingSession) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.QuizID == "" {
		errs = append(errs, "'QuizID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateQuizSubmissionStartQuizTakingSession) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
