package models

import (
	"time"
)

type Appointment struct {
	ID      int64     `json:"id" url:"id,omitempty"`             // The appointment identifier..Example: 987
	StartAt time.Time `json:"start_at" url:"start_at,omitempty"` // Start time for the appointment.Example: 2012-07-20T15:00:00-06:00
	EndAt   time.Time `json:"end_at" url:"end_at,omitempty"`     // End time for the appointment.Example: 2012-07-20T15:00:00-06:00
}

func (t *Appointment) HasErrors() error {
	return nil
}
