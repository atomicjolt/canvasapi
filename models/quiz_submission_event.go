package models

import (
	"time"
)

type QuizSubmissionEvent struct {
	CreatedAt time.Time `json:"created_at"` // a timestamp record of creation time.Example: 2014-10-08T19:29:58Z
	EventType string    `json:"event_type"` // the type of event being sent.Example: question_answered
	EventData string    `json:"event_data"` // custom contextual data for the specific event type.Example: 42
}

func (t *QuizSubmissionEvent) HasError() error {
	return nil
}
