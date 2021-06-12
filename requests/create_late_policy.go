package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// CreateLatePolicy Create a late policy. If the course already has a late policy, a
// bad_request is returned since there can only be one late policy
// per course.
// https://canvas.instructure.com/doc/api/late_policy.html
//
// Path Parameters:
// # ID (Required) ID
//
// Form Parameters:
// # LatePolicy (Optional) Whether to enable the missing submission deduction late policy.
// # LatePolicy (Optional) How many percentage points to deduct from a missing submission.
// # LatePolicy (Optional) Whether to enable the late submission deduction late policy.
// # LatePolicy (Optional) How many percentage points to deduct per the late submission interval.
// # LatePolicy (Optional) The interval for late policies.
// # LatePolicy (Optional) Whether to enable the late submission minimum percent for a late policy.
// # LatePolicy (Optional) The minimum grade a submissions can have in percentage points.
//
type CreateLatePolicy struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`

	Form struct {
		LatePolicy struct {
			MissingSubmissionDeductionEnabled   bool    `json:"missing_submission_deduction_enabled"`    //  (Optional)
			MissingSubmissionDeduction          float64 `json:"missing_submission_deduction"`            //  (Optional)
			LateSubmissionDeductionEnabled      bool    `json:"late_submission_deduction_enabled"`       //  (Optional)
			LateSubmissionDeduction             float64 `json:"late_submission_deduction"`               //  (Optional)
			LateSubmissionInterval              string  `json:"late_submission_interval"`                //  (Optional)
			LateSubmissionMinimumPercentEnabled bool    `json:"late_submission_minimum_percent_enabled"` //  (Optional)
			LateSubmissionMinimumPercent        float64 `json:"late_submission_minimum_percent"`         //  (Optional)
		} `json:"late_policy"`
	} `json:"form"`
}

func (t *CreateLatePolicy) GetMethod() string {
	return "POST"
}

func (t *CreateLatePolicy) GetURLPath() string {
	path := "courses/{id}/late_policy"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *CreateLatePolicy) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateLatePolicy) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateLatePolicy) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateLatePolicy) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
