package models

import (
	"time"
)

type LatePolicy struct {
	ID                                  int64     `json:"id"`                                      // the unique identifier for the late policy.Example: 123
	CourseID                            int64     `json:"course_id"`                               // the unique identifier for the course.Example: 123
	MissingSubmissionDeductionEnabled   bool      `json:"missing_submission_deduction_enabled"`    // whether to enable missing submission deductions.Example: true
	MissingSubmissionDeduction          float64   `json:"missing_submission_deduction"`            // amount of percentage points to deduct.Example: 12.34
	LateSubmissionDeductionEnabled      bool      `json:"late_submission_deduction_enabled"`       // whether to enable late submission deductions.Example: true
	LateSubmissionDeduction             float64   `json:"late_submission_deduction"`               // amount of percentage points to deduct per late_submission_interval.Example: 12.34
	LateSubmissionInterval              string    `json:"late_submission_interval"`                // time interval for late submission deduction.Example: hour
	LateSubmissionMinimumPercentEnabled bool      `json:"late_submission_minimum_percent_enabled"` // whether to enable late submission minimum percent.Example: true
	LateSubmissionMinimumPercent        float64   `json:"late_submission_minimum_percent"`         // the minimum score a submission can receive in percentage points.Example: 12.34
	CreatedAt                           time.Time `json:"created_at"`                              // the time at which this late policy was originally created.Example: 2012-07-01T23:59:00-06:00
	UpdatedAt                           time.Time `json:"updated_at"`                              // the time at which this late policy was last modified in any way.Example: 2012-07-01T23:59:00-06:00
}

func (t *LatePolicy) HasError() error {
	return nil
}
