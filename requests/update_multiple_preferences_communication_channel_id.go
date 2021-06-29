package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdateMultiplePreferencesCommunicationChannelID Change the preferences for multiple notifications for a single communication channel at once
// https://canvas.instructure.com/doc/api/notification_preferences.html
//
// Path Parameters:
// # Path.CommunicationChannelID (Required) ID
//
// Form Parameters:
// # Form.NotificationPreferences (Required) The desired frequency for <X> notification
//
type UpdateMultiplePreferencesCommunicationChannelID struct {
	Path struct {
		CommunicationChannelID string `json:"communication_channel_id" url:"communication_channel_id,omitempty"` //  (Required)
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

func (t *UpdateMultiplePreferencesCommunicationChannelID) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateMultiplePreferencesCommunicationChannelID) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateMultiplePreferencesCommunicationChannelID) HasErrors() error {
	errs := []string{}
	if t.Path.CommunicationChannelID == "" {
		errs = append(errs, "'Path.CommunicationChannelID' is required")
	}
	if t.Form.NotificationPreferences == nil {
		errs = append(errs, "'Form.NotificationPreferences' is required")
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
	Frequency string `json:"frequency" url:"frequency,omitempty"` //  (Required)
}
