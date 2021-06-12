package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// GetAllQuizSubmissions Get a list of all submissions for this quiz. Users who can view or manage
// grades for a course will have submissions from multiple users returned. A
// user who can only submit will have only their own submissions returned. When
// a user has an in-progress submission, only that submission is returned. When
// there isn't an in-progress quiz_submission, all completed submissions,
// including previous attempts, are returned.
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
type GetAllQuizSubmissions struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		QuizID   string `json:"quiz_id"`   //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include"` //  (Optional) . Must be one of submission, quiz, user
	} `json:"query"`
}

func (t *GetAllQuizSubmissions) GetMethod() string {
	return "GET"
}

func (t *GetAllQuizSubmissions) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/submissions"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	return path
}

func (t *GetAllQuizSubmissions) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetAllQuizSubmissions) GetBody() (string, error) {
	return "", nil
}

func (t *GetAllQuizSubmissions) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.QuizID == "" {
		errs = append(errs, "'QuizID' is required")
	}
	for _, v := range t.Query.Include {
		if !string_utils.Include([]string{"submission", "quiz", "user"}, v) {
			errs = append(errs, "Include must be one of submission, quiz, user")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetAllQuizSubmissions) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
