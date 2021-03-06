package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// ListLiveAssessmentResults Returns a paginated list of live assessment results
// https://canvas.instructure.com/doc/api/live_assessments.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.AssessmentID (Required) ID
//
// Query Parameters:
// # Query.UserID (Optional) If set, restrict results to those for this user
//
type ListLiveAssessmentResults struct {
	Path struct {
		CourseID     string `json:"course_id" url:"course_id,omitempty"`         //  (Required)
		AssessmentID string `json:"assessment_id" url:"assessment_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		UserID int64 `json:"user_id" url:"user_id,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ListLiveAssessmentResults) GetMethod() string {
	return "GET"
}

func (t *ListLiveAssessmentResults) GetURLPath() string {
	path := "courses/{course_id}/live_assessments/{assessment_id}/results"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assessment_id}", fmt.Sprintf("%v", t.Path.AssessmentID))
	return path
}

func (t *ListLiveAssessmentResults) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListLiveAssessmentResults) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListLiveAssessmentResults) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListLiveAssessmentResults) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.AssessmentID == "" {
		errs = append(errs, "'Path.AssessmentID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListLiveAssessmentResults) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
