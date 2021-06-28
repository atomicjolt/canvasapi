package models

import (
	"time"
)

type Collaboration struct {
	ID                int64     `json:"id" url:"id,omitempty"`                                 // The unique identifier for the collaboration.Example: 43
	CollaborationType string    `json:"collaboration_type" url:"collaboration_type,omitempty"` // A name for the type of collaboration.Example: Microsoft Office
	DocumentID        string    `json:"document_id" url:"document_id,omitempty"`               // The collaboration document identifier for the collaboration provider.Example: oinwoenfe8w8ef_onweufe89fef
	UserID            int64     `json:"user_id" url:"user_id,omitempty"`                       // The canvas id of the user who created the collaboration.Example: 92
	ContextID         int64     `json:"context_id" url:"context_id,omitempty"`                 // The canvas id of the course or group to which the collaboration belongs.Example: 77
	ContextType       string    `json:"context_type" url:"context_type,omitempty"`             // The canvas type of the course or group to which the collaboration belongs.Example: Course
	Url               string    `json:"url" url:"url,omitempty"`                               // The LTI launch url to view collaboration..
	CreatedAt         time.Time `json:"created_at" url:"created_at,omitempty"`                 // The timestamp when the collaboration was created.Example: 2012-06-01T00:00:00-06:00
	UpdatedAt         time.Time `json:"updated_at" url:"updated_at,omitempty"`                 // The timestamp when the collaboration was last modified.Example: 2012-06-01T00:00:00-06:00
	Description       string    `json:"description" url:"description,omitempty"`               //
	Title             string    `json:"title" url:"title,omitempty"`                           //
	Type              string    `json:"type" url:"type,omitempty"`                             // Another representation of the collaboration type.Example: ExternalToolCollaboration
	UpdateUrl         string    `json:"update_url" url:"update_url,omitempty"`                 // The LTI launch url to edit the collaboration.
	UserName          string    `json:"user_name" url:"user_name,omitempty"`                   // The name of the user who owns the collaboration.Example: John Danger
}

func (t *Collaboration) HasError() error {
	return nil
}
