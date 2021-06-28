package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ListLiveAssessments Returns a paginated list of live assessments.
// https://canvas.instructure.com/doc/api/live_assessments.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type ListLiveAssessments struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListLiveAssessments) GetMethod() string {
	return "GET"
}

func (t *ListLiveAssessments) GetURLPath() string {
	path := "courses/{course_id}/live_assessments"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListLiveAssessments) GetQuery() (string, error) {
	return "", nil
}

func (t *ListLiveAssessments) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListLiveAssessments) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListLiveAssessments) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListLiveAssessments) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
