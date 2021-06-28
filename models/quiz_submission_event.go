package models

import (
	"time"
)

type QuizSubmissionEvent struct {
	CreatedAt time.Time `json:"created_at" url:"created_at,omitempty"` // a timestamp record of creation time.Example: 2014-10-08T19:29:58Z
	EventType string    `json:"event_type" url:"event_type,omitempty"` // the type of event being sent.Example: question_answered
	EventData string    `json:"event_data" url:"event_data,omitempty"` // custom contextual data for the specific event type.Example: 42
}

func (t *QuizSubmissionEvent) HasError() error {
	return nil
}
