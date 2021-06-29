package models

import (
	"time"
)

type GradeChangeEvent struct {
	ID                string                 `json:"id" url:"id,omitempty"`                                 // ID of the event..Example: e2b76430-27a5-0131-3ca1-48e0eb13f29b
	CreatedAt         time.Time              `json:"created_at" url:"created_at,omitempty"`                 // timestamp of the event.Example: 2012-07-19T15:00:00-06:00
	EventType         string                 `json:"event_type" url:"event_type,omitempty"`                 // GradeChange event type.Example: grade_change
	ExcusedAfter      bool                   `json:"excused_after" url:"excused_after,omitempty"`           // Boolean indicating whether the submission was excused after the change..Example: true
	ExcusedBefore     bool                   `json:"excused_before" url:"excused_before,omitempty"`         // Boolean indicating whether the submission was excused before the change..
	GradeAfter        string                 `json:"grade_after" url:"grade_after,omitempty"`               // The grade after the change..Example: 8
	GradeBefore       string                 `json:"grade_before" url:"grade_before,omitempty"`             // The grade before the change..Example: 8
	GradedAnonymously bool                   `json:"graded_anonymously" url:"graded_anonymously,omitempty"` // Boolean indicating whether the student name was visible when the grade was given. Could be null if the grade change record was created before this feature existed..Example: true
	VersionNumber     string                 `json:"version_number" url:"version_number,omitempty"`         // Version Number of the grade change submission..Example: 1
	RequestID         string                 `json:"request_id" url:"request_id,omitempty"`                 // The unique request id of the request during the grade change..Example: e2b76430-27a5-0131-3ca1-48e0eb13f29b
	Links             *GradeChangeEventLinks `json:"links" url:"links,omitempty"`                           //
}

func (t *GradeChangeEvent) HasErrors() error {
	return nil
}
