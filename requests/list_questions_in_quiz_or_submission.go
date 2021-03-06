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

// ListQuestionsInQuizOrSubmission Returns the paginated list of QuizQuestions in this quiz.
// https://canvas.instructure.com/doc/api/quiz_questions.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.QuizID (Required) ID
//
// Query Parameters:
// # Query.QuizSubmissionID (Optional) If specified, the endpoint will return the questions that were presented
//    for that submission. This is useful if the quiz has been modified after
//    the submission was created and the latest quiz version's set of questions
//    does not match the submission's.
//    NOTE: you must specify quiz_submission_attempt as well if you specify this
//    parameter.
// # Query.QuizSubmissionAttempt (Optional) The attempt of the submission you want the questions for.
//
type ListQuestionsInQuizOrSubmission struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		QuizID   string `json:"quiz_id" url:"quiz_id,omitempty"`     //  (Required)
	} `json:"path"`

	Query struct {
		QuizSubmissionID      int64 `json:"quiz_submission_id" url:"quiz_submission_id,omitempty"`           //  (Optional)
		QuizSubmissionAttempt int64 `json:"quiz_submission_attempt" url:"quiz_submission_attempt,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ListQuestionsInQuizOrSubmission) GetMethod() string {
	return "GET"
}

func (t *ListQuestionsInQuizOrSubmission) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/questions"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	return path
}

func (t *ListQuestionsInQuizOrSubmission) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListQuestionsInQuizOrSubmission) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListQuestionsInQuizOrSubmission) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListQuestionsInQuizOrSubmission) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.QuizID == "" {
		errs = append(errs, "'Path.QuizID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListQuestionsInQuizOrSubmission) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.QuizQuestion, *canvasapi.PagedResource, error) {
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
	ret := []*models.QuizQuestion{}
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
