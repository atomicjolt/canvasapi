package models

import (
	"time"
)

type Appointment struct {
	ID      int64     `json:"id"`       // The appointment identifier..Example: 987
	StartAt time.Time `json:"start_at"` // Start time for the appointment.Example: 2012-07-20T15:00:00-06:00
	EndAt   time.Time `json:"end_at"`   // End time for the appointment.Example: 2012-07-20T15:00:00-06:00
}

func (t *Appointment) HasError() error {
	return nil
}
