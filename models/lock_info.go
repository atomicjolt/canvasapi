package models

import (
	"time"
)

type LockInfo struct {
	AssetString    string    `json:"asset_string" url:"asset_string,omitempty"`       // Asset string for the object causing the lock.Example: assignment_4
	UnlockAt       time.Time `json:"unlock_at" url:"unlock_at,omitempty"`             // (Optional) Time at which this was/will be unlocked. Must be before the due date..Example: 2013-01-01T00:00:00-06:00
	LockAt         time.Time `json:"lock_at" url:"lock_at,omitempty"`                 // (Optional) Time at which this was/will be locked. Must be after the due date..Example: 2013-02-01T00:00:00-06:00
	ContextModule  string    `json:"context_module" url:"context_module,omitempty"`   // (Optional) Context module causing the lock..Example: {}
	ManuallyLocked bool      `json:"manually_locked" url:"manually_locked,omitempty"` // Example: true
}

func (t *LockInfo) HasErrors() error {
	return nil
}
