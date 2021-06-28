package models

import (
	"time"
)

type SectionAssignmentOverrideAttributes struct {
	OverrideTitle string    `json:"override_title" url:"override_title,omitempty"` // The title for the assignment override.Example: some section override
	DueAt         time.Time `json:"due_at" url:"due_at,omitempty"`                 // the due date for the assignment. returns null if not present. NOTE: If this assignment has assignment overrides, this field will be the due date as it applies to the user requesting information from the API..Example: 2012-07-01T23:59:00-06:00
	UnlockAt      time.Time `json:"unlock_at" url:"unlock_at,omitempty"`           // (Optional) Time at which this was/will be unlocked..Example: 2013-01-01T00:00:00-06:00
	LockAt        time.Time `json:"lock_at" url:"lock_at,omitempty"`               // (Optional) Time at which this was/will be locked..Example: 2013-02-01T00:00:00-06:00
}

func (t *SectionAssignmentOverrideAttributes) HasError() error {
	return nil
}
