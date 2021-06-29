package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// SubmitCapturedEvents Store a set of events which were captured during a quiz taking session.
//
// On success, the response will be 204 No Content with an empty body.
// https://canvas.instructure.com/doc/api/quiz_submission_events.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.QuizID (Required) ID
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.QuizSubmissionEvents (Required) The submission events to be recorded
//
type SubmitCapturedEvents struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		QuizID   string `json:"quiz_id" url:"quiz_id,omitempty"`     //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Form struct {
		QuizSubmissionEvents []string `json:"quiz_submission_events" url:"quiz_submission_events,omitempty"` //  (Required)
	} `json:"form"`
}

func (t *SubmitCapturedEvents) GetMethod() string {
	return "POST"
}

func (t *SubmitCapturedEvents) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/submissions/{id}/events"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *SubmitCapturedEvents) GetQuery() (string, error) {
	return "", nil
}

func (t *SubmitCapturedEvents) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *SubmitCapturedEvents) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *SubmitCapturedEvents) HasErrors() error {
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
	if t.Form.QuizSubmissionEvents == nil {
		errs = append(errs, "'Form.QuizSubmissionEvents' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *SubmitCapturedEvents) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
