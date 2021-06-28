package models

import (
	"time"
)

type SubmissionComment struct {
	ID           int64         `json:"id" url:"id,omitempty"`                       // Example: 37
	AuthorID     int64         `json:"author_id" url:"author_id,omitempty"`         // Example: 134
	AuthorName   string        `json:"author_name" url:"author_name,omitempty"`     // Example: Toph Beifong
	Author       string        `json:"author" url:"author,omitempty"`               // Abbreviated user object UserDisplay (see users API)..Example: {}
	Comment      string        `json:"comment" url:"comment,omitempty"`             // Example: Well here's the thing.
	CreatedAt    time.Time     `json:"created_at" url:"created_at,omitempty"`       // Example: 2012-01-01T01:00:00Z
	EditedAt     time.Time     `json:"edited_at" url:"edited_at,omitempty"`         // Example: 2012-01-02T01:00:00Z
	MediaComment *MediaComment `json:"media_comment" url:"media_comment,omitempty"` //
}

func (t *SubmissionComment) HasError() error {
	return nil
}
