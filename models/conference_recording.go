package models

import (
	"time"
)

type ConferenceRecording struct {
	DurationMinutes int64     `json:"duration_minutes"` // Example: 0
	Title           string    `json:"title"`            // Example: course2: Test conference 3 [170]_0
	UpdatedAt       time.Time `json:"updated_at"`       // Example: 2013-12-12T16:09:33.903-07:00
	CreatedAt       time.Time `json:"created_at"`       // Example: 2013-12-12T16:09:09.960-07:00
	PlaybackUrl     string    `json:"playback_url"`     // Example: http://example.com/recording_url
}

func (t *ConferenceRecording) HasError() error {
	return nil
}
