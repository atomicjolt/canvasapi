package models

import (
	"time"
)

type LockInfo struct {
	AssetString    string    `json:"asset_string"`    // Asset string for the object causing the lock.Example: assignment_4
	UnlockAt       time.Time `json:"unlock_at"`       // (Optional) Time at which this was/will be unlocked. Must be before the due date..Example: 2013-01-01T00:00:00-06:00
	LockAt         time.Time `json:"lock_at"`         // (Optional) Time at which this was/will be locked. Must be after the due date..Example: 2013-02-01T00:00:00-06:00
	ContextModule  string    `json:"context_module"`  // (Optional) Context module causing the lock..Example: {}
	ManuallyLocked bool      `json:"manually_locked"` // Example: true
}

func (t *LockInfo) HasError() error {
	return nil
}
