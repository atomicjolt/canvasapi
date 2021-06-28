package models

import (
	"fmt"

	"github.com/atomicjolt/string_utils"
)

type NotificationPreference struct {
	Href         string `json:"href" url:"href,omitempty"`                 // Example: https://canvas.instructure.com/users/1/communication_channels/email/student@example.edu/notification_preferences/new_announcement
	Notification string `json:"notification" url:"notification,omitempty"` // The notification this preference belongs to.Example: new_announcement
	Category     string `json:"category" url:"category,omitempty"`         // The category of that notification.Example: announcement
	Frequency    string `json:"frequency" url:"frequency,omitempty"`       // How often to send notifications to this communication channel for the given notification. Possible values are 'immediately', 'daily', 'weekly', and 'never'.Example: daily
}

func (t *NotificationPreference) HasError() error {
	var s []string
	s = []string{"immediately", "daily", "weekly", "never"}
	if t.Frequency != "" && !string_utils.Include(s, t.Frequency) {
		return fmt.Errorf("expected 'frequency' to be one of %v", s)
	}
	return nil
}
