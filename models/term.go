package models

import (
	"time"
)

type Term struct {
	ID      int64     `json:"id" url:"id,omitempty"`             // Example: 1
	Name    string    `json:"name" url:"name,omitempty"`         // Example: Default Term
	StartAt time.Time `json:"start_at" url:"start_at,omitempty"` // Example: 2012-06-01T00:00:00-06:00
	EndAt   time.Time `json:"end_at" url:"end_at,omitempty"`     //
}

func (t *Term) HasErrors() error {
	return nil
}
