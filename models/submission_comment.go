package models

import (
	"time"
)

type SubmissionComment struct {
	ID           int64         `json:"id"`            // Example: 37
	AuthorID     int64         `json:"author_id"`     // Example: 134
	AuthorName   string        `json:"author_name"`   // Example: Toph Beifong
	Author       string        `json:"author"`        // Abbreviated user object UserDisplay (see users API)..Example: {}
	Comment      string        `json:"comment"`       // Example: Well here's the thing.
	CreatedAt    time.Time     `json:"created_at"`    // Example: 2012-01-01T01:00:00Z
	EditedAt     time.Time     `json:"edited_at"`     // Example: 2012-01-02T01:00:00Z
	MediaComment *MediaComment `json:"media_comment"` //
}

func (t *SubmissionComment) HasError() error {
	return nil
}
