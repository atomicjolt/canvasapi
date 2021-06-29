package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// FetchingLatestQuizStatistics This endpoint provides statistics for all quiz versions, or for a specific
// quiz version, in which case the output is guaranteed to represent the
// _latest_ and most current version of the quiz.
//
// <b>200 OK</b> response code is returned if the request was successful.
// https://canvas.instructure.com/doc/api/quiz_statistics.html
//
// Path Parameters:
// # CourseID (Required) ID
// # QuizID (Required) ID
//
// Query Parameters:
// # AllVersions (Optional) Whether the statistics report should include all submissions attempts.
//
type FetchingLatestQuizStatistics struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		QuizID   string `json:"quiz_id" url:"quiz_id,omitempty"`     //  (Required)
	} `json:"path"`

	Query struct {
		AllVersions bool `json:"all_versions" url:"all_versions,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *FetchingLatestQuizStatistics) GetMethod() string {
	return "GET"
}

func (t *FetchingLatestQuizStatistics) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/statistics"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	return path
}

func (t *FetchingLatestQuizStatistics) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *FetchingLatestQuizStatistics) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *FetchingLatestQuizStatistics) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *FetchingLatestQuizStatistics) HasErrors() error {
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

func (t *FetchingLatestQuizStatistics) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
