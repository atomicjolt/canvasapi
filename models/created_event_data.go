package models

type CreatedEventData struct {
	Name          []string `json:"name" url:"name,omitempty"`                     // Example: , Course 1
	StartAt       []string `json:"start_at" url:"start_at,omitempty"`             // Example: , 2012-01-19T15:00:00-06:00
	ConcludeAt    []string `json:"conclude_at" url:"conclude_at,omitempty"`       // Example: , 2012-01-19T15:00:00-08:00
	IsPublic      []string `json:"is_public" url:"is_public,omitempty"`           // Example: , false
	CreatedSource string   `json:"created_source" url:"created_source,omitempty"` // The type of action that triggered the creation of the course..Example: manual|sis|api
}

func (t *CreatedEventData) HasErrors() error {
	return nil
}
