package models

import (
	"time"
)

type ContentDetails struct {
	PointsPossible  int64     `json:"points_possible" url:"points_possible,omitempty"`   // Example: 20
	DueAt           time.Time `json:"due_at" url:"due_at,omitempty"`                     // Example: 2012-12-31T06:00:00-06:00
	UnlockAt        time.Time `json:"unlock_at" url:"unlock_at,omitempty"`               // Example: 2012-12-31T06:00:00-06:00
	LockAt          time.Time `json:"lock_at" url:"lock_at,omitempty"`                   // Example: 2012-12-31T06:00:00-06:00
	LockedForUser   bool      `json:"locked_for_user" url:"locked_for_user,omitempty"`   // Example: true
	LockExplanation string    `json:"lock_explanation" url:"lock_explanation,omitempty"` // Example: This quiz is part of an unpublished module and is not available yet.
	LockInfo        *LockInfo `json:"lock_info" url:"lock_info,omitempty"`               // Example: assignment_4, 2012-12-31T06:00:00-06:00, 2012-12-31T06:00:00-06:00, {}
}

func (t *ContentDetails) HasError() error {
	return nil
}
