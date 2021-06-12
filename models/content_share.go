package models

import (
	"time"
)

type ContentShare struct {
	ID            int64          `json:"id"`             // The id of the content share for the current user.Example: 1
	Name          string         `json:"name"`           // The name of the shared content.Example: War of 1812 homework
	ContentType   string         `json:"content_type"`   // The type of content that was shared. Can be assignment, discussion_topic, page, quiz, module, or module_item..Example: assignment
	CreatedAt     time.Time      `json:"created_at"`     // The datetime the content was shared with this user..Example: 2017-05-09T10:12:00Z
	UpdatedAt     time.Time      `json:"updated_at"`     // The datetime the content was updated..Example: 2017-05-09T10:12:00Z
	UserID        int64          `json:"user_id"`        // The id of the user who sent or received the content share..Example: 1578941
	Sender        string         `json:"sender"`         // The user who shared the content. This field is provided only to receivers; it is not populated in the sender's list of sent content shares..Example: 1, Matilda Vargas, http://localhost:3000/image_url, http://localhost:3000/users/1
	Receivers     string         `json:"receivers"`      // An Array of users the content is shared with.  This field is provided only to senders; an empty array will be returned for the receiving users..Example: {'id'=>1, 'display_name'=>'Jon Snow', 'avatar_image_url'=>'http://localhost:3000/image_url2', 'html_url'=>'http://localhost:3000/users/2'}
	SourceCourse  string         `json:"source_course"`  // The course the content was originally shared from..Example: 787, History 105
	ReadState     string         `json:"read_state"`     // Whether the recipient has viewed the content share..Example: read
	ContentExport *ContentExport `json:"content_export"` // The content export record associated with this content share.Example: 42
}

func (t *ContentShare) HasError() error {
	return nil
}
