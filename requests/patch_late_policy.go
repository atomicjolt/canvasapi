package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// PatchLatePolicy Patch a late policy. No body is returned upon success.
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
type PatchLatePolicy struct {
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

func (t *PatchLatePolicy) GetMethod() string {
	return "PATCH"
}

func (t *PatchLatePolicy) GetURLPath() string {
	path := "courses/{id}/late_policy"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *PatchLatePolicy) GetQuery() (string, error) {
	return "", nil
}

func (t *PatchLatePolicy) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *PatchLatePolicy) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *PatchLatePolicy) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
