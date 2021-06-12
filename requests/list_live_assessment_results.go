package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// ListLiveAssessmentResults Returns a paginated list of live assessment results
// https://canvas.instructure.com/doc/api/live_assessments.html
//
// Path Parameters:
// # CourseID (Required) ID
// # AssessmentID (Required) ID
//
// Query Parameters:
// # UserID (Optional) If set, restrict results to those for this user
//
type ListLiveAssessmentResults struct {
	Path struct {
		CourseID     string `json:"course_id"`     //  (Required)
		AssessmentID string `json:"assessment_id"` //  (Required)
	} `json:"path"`

	Query struct {
		UserID int64 `json:"user_id"` //  (Optional)
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
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListLiveAssessmentResults) GetBody() (string, error) {
	return "", nil
}

func (t *ListLiveAssessmentResults) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.AssessmentID == "" {
		errs = append(errs, "'AssessmentID' is required")
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
