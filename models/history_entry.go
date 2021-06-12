package models

import (
	"time"
)

type HistoryEntry struct {
	AssetCode             string    `json:"asset_code"`              // The asset string for the item viewed.Example: assignment_123
	AssetName             string    `json:"asset_name"`              // The name of the item.Example: Test Assignment
	AssetIcon             string    `json:"asset_icon"`              // The icon type shown for the item. One of 'icon-announcement', 'icon-assignment', 'icon-calendar-month', 'icon-discussion', 'icon-document', 'icon-download', 'icon-gradebook', 'icon-home', 'icon-message', 'icon-module', 'icon-outcomes', 'icon-quiz', 'icon-user', 'icon-syllabus'.Example: icon-assignment
	AssetReadableCategory string    `json:"asset_readable_category"` // The associated category describing the asset_icon.Example: Assignment
	ContextType           string    `json:"context_type"`            // The type of context of the item visited. One of 'Course', 'Group', 'User', or 'Account'.Example: Course
	ContextID             int64     `json:"context_id"`              // The id of the context, if applicable.Example: 123
	ContextName           string    `json:"context_name"`            // The name of the context.Example: Something 101
	VisitedUrl            string    `json:"visited_url"`             // The URL of the item.Example: https://canvas.example.com/courses/123/assignments/456
	VisitedAt             time.Time `json:"visited_at"`              // When the page was visited.Example: 2019-08-01T19:49:47Z
	InteractionSeconds    int64     `json:"interaction_seconds"`     // The estimated time spent on the page in seconds.Example: 400
}

func (t *HistoryEntry) HasError() error {
	return nil
}
