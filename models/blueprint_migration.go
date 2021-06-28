package models

import (
	"time"
)

type BlueprintMigration struct {
	ID                 int64     `json:"id" url:"id,omitempty"`                                     // The ID of the migration..Example: 1
	TemplateID         int64     `json:"template_id" url:"template_id,omitempty"`                   // The ID of the template the migration belongs to. Only present when querying a blueprint course..Example: 2
	SubscriptionID     int64     `json:"subscription_id" url:"subscription_id,omitempty"`           // The ID of the associated course's blueprint subscription. Only present when querying a course associated with a blueprint..Example: 101
	UserID             int64     `json:"user_id" url:"user_id,omitempty"`                           // The ID of the user who queued the migration..Example: 3
	WorkflowState      string    `json:"workflow_state" url:"workflow_state,omitempty"`             // Current state of the content migration: queued, exporting, imports_queued, completed, exports_failed, imports_failed.Example: running
	CreatedAt          time.Time `json:"created_at" url:"created_at,omitempty"`                     // Time when the migration was queued.Example: 2013-08-28T23:59:00-06:00
	ExportsStartedAt   time.Time `json:"exports_started_at" url:"exports_started_at,omitempty"`     // Time when the exports begun.Example: 2013-08-28T23:59:00-06:00
	ImportsQueuedAt    time.Time `json:"imports_queued_at" url:"imports_queued_at,omitempty"`       // Time when the exports were completed and imports were queued.Example: 2013-08-28T23:59:00-06:00
	ImportsCompletedAt time.Time `json:"imports_completed_at" url:"imports_completed_at,omitempty"` // Time when the imports were completed.Example: 2013-08-28T23:59:00-06:00
	Comment            string    `json:"comment" url:"comment,omitempty"`                           // User-specified comment describing changes made in this operation.Example: Fixed spelling in question 3 of midterm exam
}

func (t *BlueprintMigration) HasError() error {
	return nil
}
