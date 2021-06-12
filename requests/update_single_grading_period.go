package requests

import (
	"fmt"
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
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`

	Form struct {
		GradingPeriods struct {
			StartDate []time.Time `json:"start_date"` //  (Required)
			EndDate   []time.Time `json:"end_date"`   //  (Required)
			Weight    []float64   `json:"weight"`     //  (Optional)
		} `json:"grading_periods"`
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

func (t *UpdateSingleGradingPeriod) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
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
