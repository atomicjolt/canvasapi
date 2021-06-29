package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// CreateLiveAssessmentResults Creates live assessment results and adds them to a live assessment
// https://canvas.instructure.com/doc/api/live_assessments.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.AssessmentID (Required) ID
//
type CreateLiveAssessmentResults struct {
	Path struct {
		CourseID     string `json:"course_id" url:"course_id,omitempty"`         //  (Required)
		AssessmentID string `json:"assessment_id" url:"assessment_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *CreateLiveAssessmentResults) GetMethod() string {
	return "POST"
}

func (t *CreateLiveAssessmentResults) GetURLPath() string {
	path := "courses/{course_id}/live_assessments/{assessment_id}/results"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assessment_id}", fmt.Sprintf("%v", t.Path.AssessmentID))
	return path
}

func (t *CreateLiveAssessmentResults) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateLiveAssessmentResults) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *CreateLiveAssessmentResults) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *CreateLiveAssessmentResults) HasErrors() error {
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

func (t *CreateLiveAssessmentResults) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
