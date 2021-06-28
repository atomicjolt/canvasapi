package models

type Bookmark struct {
	ID       int64  `json:"id" url:"id,omitempty"`             // Example: 1
	Name     string `json:"name" url:"name,omitempty"`         // Example: Biology 101
	Url      string `json:"url" url:"url,omitempty"`           // Example: /courses/1
	Position int64  `json:"position" url:"position,omitempty"` // Example: 1
	Data     string `json:"data" url:"data,omitempty"`         // Example: 1
}

func (t *Bookmark) HasError() error {
	return nil
}
