package models

type UpdatedEventData struct {
	Name       []string `json:"name" url:"name,omitempty"`               // Example: Course 1, Course 2
	StartAt    []string `json:"start_at" url:"start_at,omitempty"`       // Example: 2012-01-19T15:00:00-06:00, 2012-07-19T15:00:00-06:00
	ConcludeAt []string `json:"conclude_at" url:"conclude_at,omitempty"` // Example: 2012-01-19T15:00:00-08:00, 2012-07-19T15:00:00-08:00
	IsPublic   []string `json:"is_public" url:"is_public,omitempty"`     // Example: true, false
}

func (t *UpdatedEventData) HasErrors() error {
	return nil
}
