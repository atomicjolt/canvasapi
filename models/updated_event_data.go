package models

import (
	"time"
)

type UpdatedEventData struct {
	Name       []string    `json:"name" url:"name,omitempty"`               // Example: Course 1, Course 2
	StartAt    []time.Time `json:"start_at" url:"start_at,omitempty"`       // Example: 2012-01-19T15:00:00-06:00, 2012-07-19T15:00:00-06:00
	ConcludeAt []time.Time `json:"conclude_at" url:"conclude_at,omitempty"` // Example: 2012-01-19T15:00:00-08:00, 2012-07-19T15:00:00-08:00
	IsPublic   []bool      `json:"is_public" url:"is_public,omitempty"`     // Example: true, false
}

func (t *UpdatedEventData) HasError() error {
	return nil
}
