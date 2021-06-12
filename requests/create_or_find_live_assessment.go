package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// CreateOrFindLiveAssessment Creates or finds an existing live assessment with the given key and aligns it with
// the linked outcome
// https://canvas.instructure.com/doc/api/live_assessments.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type CreateOrFindLiveAssessment struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`
}

func (t *CreateOrFindLiveAssessment) GetMethod() string {
	return "POST"
}

func (t *CreateOrFindLiveAssessment) GetURLPath() string {
	path := "courses/{course_id}/live_assessments"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CreateOrFindLiveAssessment) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateOrFindLiveAssessment) GetBody() (string, error) {
	return "", nil
}

func (t *CreateOrFindLiveAssessment) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateOrFindLiveAssessment) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
