package models

import (
	"time"
)

type QuizAssignmentOverride struct {
	ID       int64     `json:"id" url:"id,omitempty"`               // ID of the assignment override, unless this is the base construct, in which case the 'id' field is omitted..Example: 1
	DueAt    time.Time `json:"due_at" url:"due_at,omitempty"`       // The date after which any quiz submission is considered late..Example: 2014-02-21T06:59:59Z
	UnlockAt time.Time `json:"unlock_at" url:"unlock_at,omitempty"` // Date when the quiz becomes available for taking..
	LockAt   time.Time `json:"lock_at" url:"lock_at,omitempty"`     // When the quiz will stop being available for taking. A value of null means it can always be taken..Example: 2014-02-21T06:59:59Z
	Title    string    `json:"title" url:"title,omitempty"`         // Title of the section this assignment override is for, if any..Example: Project X
	Base     bool      `json:"base" url:"base,omitempty"`           // If this property is present, it means that dates in this structure are not based on an assignment override, but are instead for all students..Example: true
}

func (t *QuizAssignmentOverride) HasError() error {
	return nil
}
