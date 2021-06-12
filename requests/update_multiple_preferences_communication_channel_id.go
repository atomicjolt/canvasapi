package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdateMultiplePreferencesCommunicationChannelID Change the preferences for multiple notifications for a single communication channel at once
// https://canvas.instructure.com/doc/api/notification_preferences.html
//
// Path Parameters:
// # CommunicationChannelID (Required) ID
//
// Form Parameters:
// # NotificationPreferences (Required) The desired frequency for <X> notification
//
type UpdateMultiplePreferencesCommunicationChannelID struct {
	Path struct {
		CommunicationChannelID string `json:"communication_channel_id"` //  (Required)
	} `json:"path"`

	Form struct {
		NotificationPreferences map[string]UpdateMultiplePreferencesCommunicationChannelIDNotificationPreferences
	} `json:"form"`
}

func (t *UpdateMultiplePreferencesCommunicationChannelID) GetMethod() string {
	return "PUT"
}

func (t *UpdateMultiplePreferencesCommunicationChannelID) GetURLPath() string {
	path := "users/self/communication_channels/{communication_channel_id}/notification_preferences"
	path = strings.ReplaceAll(path, "{communication_channel_id}", fmt.Sprintf("%v", t.Path.CommunicationChannelID))
	return path
}

func (t *UpdateMultiplePreferencesCommunicationChannelID) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateMultiplePreferencesCommunicationChannelID) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateMultiplePreferencesCommunicationChannelID) HasErrors() error {
	errs := []string{}
	if t.Path.CommunicationChannelID == "" {
		errs = append(errs, "'CommunicationChannelID' is required")
	}
	if t.Form.NotificationPreferences == nil {
		errs = append(errs, "'NotificationPreferences' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateMultiplePreferencesCommunicationChannelID) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}

type UpdateMultiplePreferencesCommunicationChannelIDNotificationPreferences struct {
	Frequency string `json:"frequency"` //  (Required)
}
