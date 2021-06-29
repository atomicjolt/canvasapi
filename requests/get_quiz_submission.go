package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// GetQuizSubmission Get the submission for this quiz for the current user.
//
// <b>200 OK</b> response code is returned if the request was successful.
// https://canvas.instructure.com/doc/api/quiz_submissions.html
//
// Path Parameters:
// # CourseID (Required) ID
// # QuizID (Required) ID
//
// Query Parameters:
// # Include (Optional) . Must be one of submission, quiz, userAssociations to include with the quiz submission.
//
type GetQuizSubmission struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		QuizID   string `json:"quiz_id" url:"quiz_id,omitempty"`     //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of submission, quiz, user
	} `json:"query"`
}

func (t *GetQuizSubmission) GetMethod() string {
	return "GET"
}

func (t *GetQuizSubmission) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/submission"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	return path
}

func (t *GetQuizSubmission) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *GetQuizSubmission) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetQuizSubmission) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetQuizSubmission) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.QuizID == "" {
		errs = append(errs, "'QuizID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"submission", "quiz", "user"}, v) {
			errs = append(errs, "Include must be one of submission, quiz, user")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetQuizSubmission) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
