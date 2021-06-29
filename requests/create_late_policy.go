package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
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
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.LatePolicy.MissingSubmissionDeductionEnabled (Optional) Whether to enable the missing submission deduction late policy.
// # Form.LatePolicy.MissingSubmissionDeduction (Optional) How many percentage points to deduct from a missing submission.
// # Form.LatePolicy.LateSubmissionDeductionEnabled (Optional) Whether to enable the late submission deduction late policy.
// # Form.LatePolicy.LateSubmissionDeduction (Optional) How many percentage points to deduct per the late submission interval.
// # Form.LatePolicy.LateSubmissionInterval (Optional) The interval for late policies.
// # Form.LatePolicy.LateSubmissionMinimumPercentEnabled (Optional) Whether to enable the late submission minimum percent for a late policy.
// # Form.LatePolicy.LateSubmissionMinimumPercent (Optional) The minimum grade a submissions can have in percentage points.
//
type CreateLatePolicy struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		LatePolicy struct {
			MissingSubmissionDeductionEnabled   bool    `json:"missing_submission_deduction_enabled" url:"missing_submission_deduction_enabled,omitempty"`       //  (Optional)
			MissingSubmissionDeduction          float64 `json:"missing_submission_deduction" url:"missing_submission_deduction,omitempty"`                       //  (Optional)
			LateSubmissionDeductionEnabled      bool    `json:"late_submission_deduction_enabled" url:"late_submission_deduction_enabled,omitempty"`             //  (Optional)
			LateSubmissionDeduction             float64 `json:"late_submission_deduction" url:"late_submission_deduction,omitempty"`                             //  (Optional)
			LateSubmissionInterval              string  `json:"late_submission_interval" url:"late_submission_interval,omitempty"`                               //  (Optional)
			LateSubmissionMinimumPercentEnabled bool    `json:"late_submission_minimum_percent_enabled" url:"late_submission_minimum_percent_enabled,omitempty"` //  (Optional)
			LateSubmissionMinimumPercent        float64 `json:"late_submission_minimum_percent" url:"late_submission_minimum_percent,omitempty"`                 //  (Optional)
		} `json:"late_policy" url:"late_policy,omitempty"`
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

func (t *CreateLatePolicy) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateLatePolicy) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateLatePolicy) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
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
