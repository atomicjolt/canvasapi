package models

import (
	"time"
)

type CreatedEventData struct {
	Name          []string    `json:"name"`           // Example: , Course 1
	StartAt       []time.Time `json:"start_at"`       // Example: , 2012-01-19T15:00:00-06:00
	ConcludeAt    []time.Time `json:"conclude_at"`    // Example: , 2012-01-19T15:00:00-08:00
	IsPublic      []bool      `json:"is_public"`      // Example: , false
	CreatedSource string      `json:"created_source"` // The type of action that triggered the creation of the course..Example: manual|sis|api
}

func (t *CreatedEventData) HasError() error {
	return nil
}
