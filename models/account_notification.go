package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type AccountNotification struct {
	Subject string    `json:"subject" url:"subject,omitempty"`   // The subject of the notifications.Example: Attention Students
	Message string    `json:"message" url:"message,omitempty"`   // The message to be sent in the notification..Example: This is a test of the notification system.
	StartAt time.Time `json:"start_at" url:"start_at,omitempty"` // When to send out the notification..Example: 2013-08-28T23:59:00-06:00
	EndAt   time.Time `json:"end_at" url:"end_at,omitempty"`     // When to expire the notification..Example: 2013-08-29T23:59:00-06:00
	Icon    string    `json:"icon" url:"icon,omitempty"`         // The icon to display with the message.  Defaults to warning..Example: information
	Roles   []string  `json:"roles" url:"roles,omitempty"`       // (Deprecated) The roles to send the notification to.  If roles is not passed it defaults to all roles.Example: StudentEnrollment
	RoleIDs []int64   `json:"role_ids" url:"role_ids,omitempty"` // The roles to send the notification to.  If roles is not passed it defaults to all roles.Example: 1
}

func (t *AccountNotification) HasError() error {
	var s []string
	s = []string{"warning", "information", "question", "error", "calendar"}
	if t.Icon != "" && !string_utils.Include(s, t.Icon) {
		return fmt.Errorf("expected 'icon' to be one of %v", s)
	}
	return nil
}
