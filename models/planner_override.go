package models

import (
	"time"
)

type PlannerOverride struct {
	ID             int64     `json:"id" url:"id,omitempty"`                           // The ID of the planner override.Example: 234
	PlannableType  string    `json:"plannable_type" url:"plannable_type,omitempty"`   // The type of the associated object for the planner override.Example: Assignment
	PlannableID    int64     `json:"plannable_id" url:"plannable_id,omitempty"`       // The id of the associated object for the planner override.Example: 1578941
	UserID         int64     `json:"user_id" url:"user_id,omitempty"`                 // The id of the associated user for the planner override.Example: 1578941
	AssignmentID   int64     `json:"assignment_id" url:"assignment_id,omitempty"`     // The id of the plannable's associated assignment, if it has one.Example: 1578941
	WorkflowState  string    `json:"workflow_state" url:"workflow_state,omitempty"`   // The current published state of the item, synced with the associated object.Example: published
	MarkedComplete bool      `json:"marked_complete" url:"marked_complete,omitempty"` // Controls whether or not the associated plannable item is marked complete on the planner.
	Dismissed      bool      `json:"dismissed" url:"dismissed,omitempty"`             // Controls whether or not the associated plannable item shows up in the opportunities list.
	CreatedAt      time.Time `json:"created_at" url:"created_at,omitempty"`           // The datetime of when the planner override was created.Example: 2017-05-09T10:12:00Z
	UpdatedAt      time.Time `json:"updated_at" url:"updated_at,omitempty"`           // The datetime of when the planner override was updated.Example: 2017-05-09T10:12:00Z
	DeletedAt      time.Time `json:"deleted_at" url:"deleted_at,omitempty"`           // The datetime of when the planner override was deleted, if applicable.Example: 2017-05-15T12:12:00Z
}

func (t *PlannerOverride) HasErrors() error {
	return nil
}
