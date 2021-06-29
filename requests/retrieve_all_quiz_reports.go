package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// RetrieveAllQuizReports Returns a list of all available reports.
// https://canvas.instructure.com/doc/api/quiz_reports.html
//
// Path Parameters:
// # CourseID (Required) ID
// # QuizID (Required) ID
//
// Query Parameters:
// # IncludesAllVersions (Optional) Whether to retrieve reports that consider all the submissions or only
//    the most recent. Defaults to false, ignored for item_analysis reports.
//
type RetrieveAllQuizReports struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		QuizID   string `json:"quiz_id" url:"quiz_id,omitempty"`     //  (Required)
	} `json:"path"`

	Query struct {
		IncludesAllVersions bool `json:"includes_all_versions" url:"includes_all_versions,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *RetrieveAllQuizReports) GetMethod() string {
	return "GET"
}

func (t *RetrieveAllQuizReports) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/reports"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	return path
}

func (t *RetrieveAllQuizReports) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *RetrieveAllQuizReports) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *RetrieveAllQuizReports) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *RetrieveAllQuizReports) HasErrors() error {
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

func (t *RetrieveAllQuizReports) Do(c *canvasapi.Canvas) ([]*models.QuizReport, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.QuizReport{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
