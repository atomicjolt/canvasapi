package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetSingleGradingPeriod Returns the grading period with the given id
// https://canvas.instructure.com/doc/api/grading_periods.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
type GetSingleGradingPeriod struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`
}

func (t *GetSingleGradingPeriod) GetMethod() string {
	return "GET"
}

func (t *GetSingleGradingPeriod) GetURLPath() string {
	path := "courses/{course_id}/grading_periods/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetSingleGradingPeriod) GetQuery() (string, error) {
	return "", nil
}

func (t *GetSingleGradingPeriod) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSingleGradingPeriod) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSingleGradingPeriod) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleGradingPeriod) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
