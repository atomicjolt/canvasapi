package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetAvailableQuizIpFilters Get a list of available IP filters for this Quiz.
//
// <b>200 OK</b> response code is returned if the request was successful.
// https://canvas.instructure.com/doc/api/quiz_ip_filters.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.QuizID (Required) ID
//
type GetAvailableQuizIpFilters struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		QuizID   string `json:"quiz_id" url:"quiz_id,omitempty"`     //  (Required)
	} `json:"path"`
}

func (t *GetAvailableQuizIpFilters) GetMethod() string {
	return "GET"
}

func (t *GetAvailableQuizIpFilters) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/ip_filters"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	return path
}

func (t *GetAvailableQuizIpFilters) GetQuery() (string, error) {
	return "", nil
}

func (t *GetAvailableQuizIpFilters) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetAvailableQuizIpFilters) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetAvailableQuizIpFilters) HasErrors() error {
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

func (t *GetAvailableQuizIpFilters) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
