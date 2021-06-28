package models

import (
	"time"
)

type File struct {
	Size        int64     `json:"size" url:"size,omitempty"`                 // Example: 4
	Contenttype string    `json:"content_type" url:"content_type,omitempty"` // Example: text/plain
	Url         string    `json:"url" url:"url,omitempty"`                   // Example: http://www.example.com/files/569/download?download_frd=1&verifier=c6HdZmxOZa0Fiin2cbvZeI8I5ry7yqD7RChQzb6P
	ID          int64     `json:"id" url:"id,omitempty"`                     // Example: 569
	DisplayName string    `json:"display_name" url:"display_name,omitempty"` // Example: file.txt
	CreatedAt   time.Time `json:"created_at" url:"created_at,omitempty"`     // Example: 2012-07-06T14:58:50Z
	UpdatedAt   time.Time `json:"updated_at" url:"updated_at,omitempty"`     // Example: 2012-07-06T14:58:50Z
}

func (t *File) HasError() error {
	return nil
}
