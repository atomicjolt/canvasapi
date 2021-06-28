package models

import (
	"time"
)

type AssignmentOverride struct {
	ID              int64     `json:"id" url:"id,omitempty"`                               // the ID of the assignment override.Example: 4
	AssignmentID    int64     `json:"assignment_id" url:"assignment_id,omitempty"`         // the ID of the assignment the override applies to.Example: 123
	StudentIDs      []int64   `json:"student_ids" url:"student_ids,omitempty"`             // the IDs of the override's target students (present if the override targets an ad-hoc set of students).Example: 1, 2, 3
	GroupID         int64     `json:"group_id" url:"group_id,omitempty"`                   // the ID of the override's target group (present if the override targets a group and the assignment is a group assignment).Example: 2
	CourseSectionID int64     `json:"course_section_id" url:"course_section_id,omitempty"` // the ID of the overrides's target section (present if the override targets a section).Example: 1
	Title           string    `json:"title" url:"title,omitempty"`                         // the title of the override.Example: an assignment override
	DueAt           time.Time `json:"due_at" url:"due_at,omitempty"`                       // the overridden due at (present if due_at is overridden).Example: 2012-07-01T23:59:00-06:00
	AllDay          bool      `json:"all_day" url:"all_day,omitempty"`                     // the overridden all day flag (present if due_at is overridden).Example: true
	AllDayDate      time.Time `json:"all_day_date" url:"all_day_date,omitempty"`           // the overridden all day date (present if due_at is overridden).Example: 2012-07-01
	UnlockAt        time.Time `json:"unlock_at" url:"unlock_at,omitempty"`                 // the overridden unlock at (present if unlock_at is overridden).Example: 2012-07-01T23:59:00-06:00
	LockAt          time.Time `json:"lock_at" url:"lock_at,omitempty"`                     // the overridden lock at, if any (present if lock_at is overridden).Example: 2012-07-01T23:59:00-06:00
}

func (t *AssignmentOverride) HasError() error {
	return nil
}
