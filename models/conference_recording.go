package models

import (
	"time"
)

type ConferenceRecording struct {
	DurationMinutes int64     `json:"duration_minutes" url:"duration_minutes,omitempty"` // Example: 0
	Title           string    `json:"title" url:"title,omitempty"`                       // Example: course2: Test conference 3 [170]_0
	UpdatedAt       time.Time `json:"updated_at" url:"updated_at,omitempty"`             // Example: 2013-12-12T16:09:33.903-07:00
	CreatedAt       time.Time `json:"created_at" url:"created_at,omitempty"`             // Example: 2013-12-12T16:09:09.960-07:00
	PlaybackUrl     string    `json:"playback_url" url:"playback_url,omitempty"`         // Example: http://example.com/recording_url
}

func (t *ConferenceRecording) HasError() error {
	return nil
}
