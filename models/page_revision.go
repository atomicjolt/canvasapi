package models

import (
	"time"
)

type PageRevision struct {
	RevisionID int64     `json:"revision_id"` // an identifier for this revision of the page.Example: 7
	UpdatedAt  time.Time `json:"updated_at"`  // the time when this revision was saved.Example: 2012-08-07T11:23:58-06:00
	Latest     bool      `json:"latest"`      // whether this is the latest revision or not.Example: true
	EditedBy   *User     `json:"edited_by"`   // the User who saved this revision, if applicable (this may not be present if the page was imported from another system).
	Url        string    `json:"url"`         // the following fields are not included in the index action and may be omitted from the show action via summary=1 the historic url of the page.Example: old-page-title
	Title      string    `json:"title"`       // the historic page title.Example: Old Page Title
	Body       string    `json:"body"`        // the historic page contents.Example: <p>Old Page Content</p>
}

func (t *PageRevision) HasError() error {
	return nil
}
