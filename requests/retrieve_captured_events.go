package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// RetrieveCapturedEvents Retrieve the set of events captured during a specific submission attempt.
// https://canvas.instructure.com/doc/api/quiz_submission_events.html
//
// Path Parameters:
// # CourseID (Required) ID
// # QuizID (Required) ID
// # ID (Required) ID
//
// Query Parameters:
// # Attempt (Optional) The specific submission attempt to look up the events for. If unspecified,
//    the latest attempt will be used.
//
type RetrieveCapturedEvents struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		QuizID   string `json:"quiz_id"`   //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`

	Query struct {
		Attempt int64 `json:"attempt"` //  (Optional)
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
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *RetrieveCapturedEvents) GetBody() (string, error) {
	return "", nil
}

func (t *RetrieveCapturedEvents) HasErrors() error {
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
