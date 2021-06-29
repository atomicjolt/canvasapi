package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// RetrieveCapturedEvents Retrieve the set of events captured during a specific submission attempt.
// https://canvas.instructure.com/doc/api/quiz_submission_events.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.QuizID (Required) ID
// # Path.ID (Required) ID
//
// Query Parameters:
// # Query.Attempt (Optional) The specific submission attempt to look up the events for. If unspecified,
//    the latest attempt will be used.
//
type RetrieveCapturedEvents struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		QuizID   string `json:"quiz_id" url:"quiz_id,omitempty"`     //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Query struct {
		Attempt int64 `json:"attempt" url:"attempt,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *RetrieveCapturedEvents) GetMethod() string {
	return "GET"
}

func (t *RetrieveCapturedEvents) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/submissions/{id}/events"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *RetrieveCapturedEvents) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *RetrieveCapturedEvents) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *RetrieveCapturedEvents) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *RetrieveCapturedEvents) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.QuizID == "" {
		errs = append(errs, "'Path.QuizID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *RetrieveCapturedEvents) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
