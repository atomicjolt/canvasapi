package models

import (
	"time"
)

type Term struct {
	ID      int64     `json:"id"`       // Example: 1
	Name    string    `json:"name"`     // Example: Default Term
	StartAt time.Time `json:"start_at"` // Example: 2012-06-01T00:00:00-06:00
	EndAt   time.Time `json:"end_at"`   //
}

func (t *Term) HasError() error {
	return nil
}
