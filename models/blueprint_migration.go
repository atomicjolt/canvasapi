package models

import (
	"time"
)

type BlueprintMigration struct {
	ID                 int64     `json:"id"`                   // The ID of the migration..Example: 1
	TemplateID         int64     `json:"template_id"`          // The ID of the template the migration belongs to. Only present when querying a blueprint course..Example: 2
	SubscriptionID     int64     `json:"subscription_id"`      // The ID of the associated course's blueprint subscription. Only present when querying a course associated with a blueprint..Example: 101
	UserID             int64     `json:"user_id"`              // The ID of the user who queued the migration..Example: 3
	WorkflowState      string    `json:"workflow_state"`       // Current state of the content migration: queued, exporting, imports_queued, completed, exports_failed, imports_failed.Example: running
	CreatedAt          time.Time `json:"created_at"`           // Time when the migration was queued.Example: 2013-08-28T23:59:00-06:00
	ExportsStartedAt   time.Time `json:"exports_started_at"`   // Time when the exports begun.Example: 2013-08-28T23:59:00-06:00
	ImportsQueuedAt    time.Time `json:"imports_queued_at"`    // Time when the exports were completed and imports were queued.Example: 2013-08-28T23:59:00-06:00
	ImportsCompletedAt time.Time `json:"imports_completed_at"` // Time when the imports were completed.Example: 2013-08-28T23:59:00-06:00
	Comment            string    `json:"comment"`              // User-specified comment describing changes made in this operation.Example: Fixed spelling in question 3 of midterm exam
}

func (t *BlueprintMigration) HasError() error {
	return nil
}
