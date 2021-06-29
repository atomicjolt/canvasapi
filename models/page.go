package models

import (
	"time"
)

type Page struct {
	Url              string    `json:"url" url:"url,omitempty"`                               // the unique locator for the page.Example: my-page-title
	Title            string    `json:"title" url:"title,omitempty"`                           // the title of the page.Example: My Page Title
	CreatedAt        time.Time `json:"created_at" url:"created_at,omitempty"`                 // the creation date for the page.Example: 2012-08-06T16:46:33-06:00
	UpdatedAt        time.Time `json:"updated_at" url:"updated_at,omitempty"`                 // the date the page was last updated.Example: 2012-08-08T14:25:20-06:00
	HideFromStudents bool      `json:"hide_from_students" url:"hide_from_students,omitempty"` // (DEPRECATED) whether this page is hidden from students (note: this is always reflected as the inverse of the published value).
	EditingRoles     string    `json:"editing_roles" url:"editing_roles,omitempty"`           // roles allowed to edit the page; comma-separated list comprising a combination of 'teachers', 'students', 'members', and/or 'public' if not supplied, course defaults are used.Example: teachers,students
	LastEditedBy     *User     `json:"last_edited_by" url:"last_edited_by,omitempty"`         // the User who last edited the page (this may not be present if the page was imported from another system).
	Body             string    `json:"body" url:"body,omitempty"`                             // the page content, in HTML (present when requesting a single page; omitted when listing pages).Example: <p>Page Content</p>
	Published        bool      `json:"published" url:"published,omitempty"`                   // whether the page is published (true) or draft state (false)..Example: true
	FrontPage        bool      `json:"front_page" url:"front_page,omitempty"`                 // whether this page is the front page for the wiki.
	LockedForUser    bool      `json:"locked_for_user" url:"locked_for_user,omitempty"`       // Whether or not this is locked for the user..
	LockInfo         *LockInfo `json:"lock_info" url:"lock_info,omitempty"`                   // (Optional) Information for the user about the lock. Present when locked_for_user is true..
	LockExplanation  string    `json:"lock_explanation" url:"lock_explanation,omitempty"`     // (Optional) An explanation of why this is locked for the user. Present when locked_for_user is true..Example: This page is locked until September 1 at 12:00am
}

func (t *Page) HasErrors() error {
	return nil
}
