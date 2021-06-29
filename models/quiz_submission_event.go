package models

import (
	"time"
)

type QuizSubmissionEvent struct {
	CreatedAt time.Time                `json:"created_at" url:"created_at,omitempty"` // a timestamp record of creation time.Example: 2014-10-08T19:29:58Z
	EventType string                   `json:"event_type" url:"event_type,omitempty"` // the type of event being sent.Example: question_answered
	EventData map[string](interface{}) `json:"event_data" url:"event_data,omitempty"` // custom contextual data for the specific event type.Example: 42
}

func (t *QuizSubmissionEvent) HasErrors() error {
	return nil
}
