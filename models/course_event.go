package models

import (
	"time"
)

type CourseEvent struct {
	ID          string           `json:"id" url:"id,omitempty"`                     // ID of the event..Example: e2b76430-27a5-0131-3ca1-48e0eb13f29b
	CreatedAt   time.Time        `json:"created_at" url:"created_at,omitempty"`     // timestamp of the event.Example: 2012-07-19T15:00:00-06:00
	EventType   string           `json:"event_type" url:"event_type,omitempty"`     // Course event type The event type defines the type and schema of the event_data object..Example: updated
	EventData   string           `json:"event_data" url:"event_data,omitempty"`     // Course event data depending on the event type.  This will return an object containing the relevant event data.  An updated event type will return an UpdatedEventData object..Example: {}
	EventSource string           `json:"event_source" url:"event_source,omitempty"` // Course event source depending on the event type.  This will return a string containing the source of the event..Example: manual|sis|api
	Links       *CourseEventLink `json:"links" url:"links,omitempty"`               // Jsonapi.org links.Example: 12345, 12345, e2b76430-27a5-0131-3ca1-48e0eb13f29b
}

func (t *CourseEvent) HasError() error {
	return nil
}
