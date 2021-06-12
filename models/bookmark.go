package models

type Bookmark struct {
	ID       int64  `json:"id"`       // Example: 1
	Name     string `json:"name"`     // Example: Biology 101
	Url      string `json:"url"`      // Example: /courses/1
	Position int64  `json:"position"` // Example: 1
	Data     string `json:"data"`     // Example: 1
}

func (t *Bookmark) HasError() error {
	return nil
}
