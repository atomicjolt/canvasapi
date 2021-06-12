package models

import (
	"time"
)

type Collaboration struct {
	ID                int64     `json:"id"`                 // The unique identifier for the collaboration.Example: 43
	CollaborationType string    `json:"collaboration_type"` // A name for the type of collaboration.Example: Microsoft Office
	DocumentID        string    `json:"document_id"`        // The collaboration document identifier for the collaboration provider.Example: oinwoenfe8w8ef_onweufe89fef
	UserID            int64     `json:"user_id"`            // The canvas id of the user who created the collaboration.Example: 92
	ContextID         int64     `json:"context_id"`         // The canvas id of the course or group to which the collaboration belongs.Example: 77
	ContextType       string    `json:"context_type"`       // The canvas type of the course or group to which the collaboration belongs.Example: Course
	Url               string    `json:"url"`                // The LTI launch url to view collaboration..
	CreatedAt         time.Time `json:"created_at"`         // The timestamp when the collaboration was created.Example: 2012-06-01T00:00:00-06:00
	UpdatedAt         time.Time `json:"updated_at"`         // The timestamp when the collaboration was last modified.Example: 2012-06-01T00:00:00-06:00
	Description       string    `json:"description"`        //
	Title             string    `json:"title"`              //
	Type              string    `json:"type"`               // Another representation of the collaboration type.Example: ExternalToolCollaboration
	UpdateUrl         string    `json:"update_url"`         // The LTI launch url to edit the collaboration.
	UserName          string    `json:"user_name"`          // The name of the user who owns the collaboration.Example: John Danger
}

func (t *Collaboration) HasError() error {
	return nil
}
