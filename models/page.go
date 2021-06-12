package models

import (
	"time"
)

type Page struct {
	Url              string    `json:"url"`                // the unique locator for the page.Example: my-page-title
	Title            string    `json:"title"`              // the title of the page.Example: My Page Title
	CreatedAt        time.Time `json:"created_at"`         // the creation date for the page.Example: 2012-08-06T16:46:33-06:00
	UpdatedAt        time.Time `json:"updated_at"`         // the date the page was last updated.Example: 2012-08-08T14:25:20-06:00
	HideFromStudents bool      `json:"hide_from_students"` // (DEPRECATED) whether this page is hidden from students (note: this is always reflected as the inverse of the published value).
	EditingRoles     string    `json:"editing_roles"`      // roles allowed to edit the page; comma-separated list comprising a combination of 'teachers', 'students', 'members', and/or 'public' if not supplied, course defaults are used.Example: teachers,students
	LastEditedBy     *User     `json:"last_edited_by"`     // the User who last edited the page (this may not be present if the page was imported from another system).
	Body             string    `json:"body"`               // the page content, in HTML (present when requesting a single page; omitted when listing pages).Example: <p>Page Content</p>
	Published        bool      `json:"published"`          // whether the page is published (true) or draft state (false)..Example: true
	FrontPage        bool      `json:"front_page"`         // whether this page is the front page for the wiki.
	LockedForUser    bool      `json:"locked_for_user"`    // Whether or not this is locked for the user..
	LockInfo         *LockInfo `json:"lock_info"`          // (Optional) Information for the user about the lock. Present when locked_for_user is true..
	LockExplanation  string    `json:"lock_explanation"`   // (Optional) An explanation of why this is locked for the user. Present when locked_for_user is true..Example: This page is locked until September 1 at 12:00am
}

func (t *Page) HasError() error {
	return nil
}
