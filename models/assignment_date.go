package models

import (
	"time"
)

type AssignmentDate struct {
	ID       int64     `json:"id" url:"id,omitempty"`               // (Optional, missing if 'base' is present) id of the assignment override this date represents.Example: 1
	Base     bool      `json:"base" url:"base,omitempty"`           // (Optional, present if 'id' is missing) whether this date represents the assignment's or quiz's default due date.Example: true
	Title    string    `json:"title" url:"title,omitempty"`         // Example: Summer Session
	DueAt    time.Time `json:"due_at" url:"due_at,omitempty"`       // The due date for the assignment. Must be between the unlock date and the lock date if there are lock dates.Example: 2013-08-28T23:59:00-06:00
	UnlockAt time.Time `json:"unlock_at" url:"unlock_at,omitempty"` // The unlock date for the assignment. Must be before the due date if there is a due date..Example: 2013-08-01T00:00:00-06:00
	LockAt   time.Time `json:"lock_at" url:"lock_at,omitempty"`     // The lock date for the assignment. Must be after the due date if there is a due date..Example: 2013-08-31T23:59:00-06:00
}

func (t *AssignmentDate) HasErrors() error {
	return nil
}
