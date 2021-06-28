package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdateSingleGradingPeriod Update an existing grading period.
// https://canvas.instructure.com/doc/api/grading_periods.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # GradingPeriods (Required) The date the grading period starts.
// # GradingPeriods (Required) no description
// # GradingPeriods (Optional) A weight value that contributes to the overall weight of a grading period set which is used to calculate how much assignments in this period contribute to the total grade
//
type UpdateSingleGradingPeriod struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Form struct {
		GradingPeriods struct {
			StartDate []time.Time `json:"start_date" url:"start_date,omitempty"` //  (Required)
			EndDate   []time.Time `json:"end_date" url:"end_date,omitempty"`     //  (Required)
			Weight    []float64   `json:"weight" url:"weight,omitempty"`         //  (Optional)
		} `json:"grading_periods" url:"grading_periods,omitempty"`
	} `json:"form"`
}

func (t *UpdateSingleGradingPeriod) GetMethod() string {
	return "PUT"
}

func (t *UpdateSingleGradingPeriod) GetURLPath() string {
	path := "courses/{course_id}/grading_periods/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateSingleGradingPeriod) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateSingleGradingPeriod) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateSingleGradingPeriod) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateSingleGradingPeriod) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Form.GradingPeriods.StartDate == nil {
		errs = append(errs, "'GradingPeriods' is required")
	}
	if t.Form.GradingPeriods.EndDate == nil {
		errs = append(errs, "'GradingPeriods' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateSingleGradingPeriod) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
