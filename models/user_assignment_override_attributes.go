package models

import (
	"time"
)

type UserAssignmentOverrideAttributes struct {
	ID       int64                `json:"id"`        // The unique Canvas identifier for the assignment override.Example: 218
	Title    string               `json:"title"`     // The title of the assignment override..Example: Override title
	DueAt    time.Time            `json:"due_at"`    // The time at which this assignment is due.Example: 2013-01-01T00:00:00-06:00
	UnlockAt time.Time            `json:"unlock_at"` // (Optional) Time at which this was/will be unlocked..Example: 2013-01-01T00:00:00-06:00
	LockAt   time.Time            `json:"lock_at"`   // (Optional) Time at which this was/will be locked..Example: 2013-02-01T00:00:00-06:00
	Students []*StudentAttributes `json:"students"`  // Includes attributes of a student for convenience. For more details see Users API..
}

func (t *UserAssignmentOverrideAttributes) HasError() error {
	return nil
}
