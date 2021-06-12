package models

import (
	"fmt"

	"github.com/atomicjolt/string_utils"
)

type NotificationPreference struct {
	Href         string `json:"href"`         // Example: https://canvas.instructure.com/users/1/communication_channels/email/student@example.edu/notification_preferences/new_announcement
	Notification string `json:"notification"` // The notification this preference belongs to.Example: new_announcement
	Category     string `json:"category"`     // The category of that notification.Example: announcement
	Frequency    string `json:"frequency"`    // How often to send notifications to this communication channel for the given notification. Possible values are 'immediately', 'daily', 'weekly', and 'never'.Example: daily
}

func (t *NotificationPreference) HasError() error {
	var s []string
	s = []string{"immediately", "daily", "weekly", "never"}
	if !string_utils.Include(s, t.Frequency) {
		return fmt.Errorf("expected 'frequency' to be one of %v", s)
	}
	return nil
}
