package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// GetSingleQuizSubmission Get a single quiz submission.
//
// <b>200 OK</b> response code is returned if the request was successful.
// https://canvas.instructure.com/doc/api/quiz_submissions.html
//
// Path Parameters:
// # CourseID (Required) ID
// # QuizID (Required) ID
// # ID (Required) ID
//
// Query Parameters:
// # Include (Optional) . Must be one of submission, quiz, userAssociations to include with the quiz submission.
//
type GetSingleQuizSubmission struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		QuizID   string `json:"quiz_id"`   //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include"` //  (Optional) . Must be one of submission, quiz, user
	} `json:"query"`
}

func (t *GetSingleQuizSubmission) GetMethod() string {
	return "GET"
}

func (t *GetSingleQuizSubmission) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/submissions/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetSingleQuizSubmission) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetSingleQuizSubmission) GetBody() (string, error) {
	return "", nil
}

func (t *GetSingleQuizSubmission) HasErrors() error {
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

func (t *GetSingleQuizSubmission) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
