package models

import (
	"time"
)

type LatePolicy struct {
	ID                                  int64     `json:"id" url:"id,omitempty"`                                                                           // the unique identifier for the late policy.Example: 123
	CourseID                            int64     `json:"course_id" url:"course_id,omitempty"`                                                             // the unique identifier for the course.Example: 123
	MissingSubmissionDeductionEnabled   bool      `json:"missing_submission_deduction_enabled" url:"missing_submission_deduction_enabled,omitempty"`       // whether to enable missing submission deductions.Example: true
	MissingSubmissionDeduction          float64   `json:"missing_submission_deduction" url:"missing_submission_deduction,omitempty"`                       // amount of percentage points to deduct.Example: 12.34
	LateSubmissionDeductionEnabled      bool      `json:"late_submission_deduction_enabled" url:"late_submission_deduction_enabled,omitempty"`             // whether to enable late submission deductions.Example: true
	LateSubmissionDeduction             float64   `json:"late_submission_deduction" url:"late_submission_deduction,omitempty"`                             // amount of percentage points to deduct per late_submission_interval.Example: 12.34
	LateSubmissionInterval              string    `json:"late_submission_interval" url:"late_submission_interval,omitempty"`                               // time interval for late submission deduction.Example: hour
	LateSubmissionMinimumPercentEnabled bool      `json:"late_submission_minimum_percent_enabled" url:"late_submission_minimum_percent_enabled,omitempty"` // whether to enable late submission minimum percent.Example: true
	LateSubmissionMinimumPercent        float64   `json:"late_submission_minimum_percent" url:"late_submission_minimum_percent,omitempty"`                 // the minimum score a submission can receive in percentage points.Example: 12.34
	CreatedAt                           time.Time `json:"created_at" url:"created_at,omitempty"`                                                           // the time at which this late policy was originally created.Example: 2012-07-01T23:59:00-06:00
	UpdatedAt                           time.Time `json:"updated_at" url:"updated_at,omitempty"`                                                           // the time at which this late policy was last modified in any way.Example: 2012-07-01T23:59:00-06:00
}

func (t *LatePolicy) HasError() error {
	return nil
}
