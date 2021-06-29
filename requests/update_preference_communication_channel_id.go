package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdatePreferenceCommunicationChannelID Change the preference for a single notification for a single communication channel
// https://canvas.instructure.com/doc/api/notification_preferences.html
//
// Path Parameters:
// # Path.CommunicationChannelID (Required) ID
// # Path.Notification (Required) ID
//
// Form Parameters:
// # Form.NotificationPreferences.Frequency (Required) The desired frequency for this notification
//
type UpdatePreferenceCommunicationChannelID struct {
	Path struct {
		CommunicationChannelID string `json:"communication_channel_id" url:"communication_channel_id,omitempty"` //  (Required)
		Notification           string `json:"notification" url:"notification,omitempty"`                         //  (Required)
	} `json:"path"`

	Form struct {
		NotificationPreferences struct {
			Frequency string `json:"frequency" url:"frequency,omitempty"` //  (Required)
		} `json:"notification_preferences" url:"notification_preferences,omitempty"`
	} `json:"form"`
}

func (t *UpdatePreferenceCommunicationChannelID) GetMethod() string {
	return "PUT"
}

func (t *UpdatePreferenceCommunicationChannelID) GetURLPath() string {
	path := "users/self/communication_channels/{communication_channel_id}/notification_preferences/{notification}"
	path = strings.ReplaceAll(path, "{communication_channel_id}", fmt.Sprintf("%v", t.Path.CommunicationChannelID))
	path = strings.ReplaceAll(path, "{notification}", fmt.Sprintf("%v", t.Path.Notification))
	return path
}

func (t *UpdatePreferenceCommunicationChannelID) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdatePreferenceCommunicationChannelID) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdatePreferenceCommunicationChannelID) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdatePreferenceCommunicationChannelID) HasErrors() error {
	errs := []string{}
	if t.Path.CommunicationChannelID == "" {
		errs = append(errs, "'Path.CommunicationChannelID' is required")
	}
	if t.Path.Notification == "" {
		errs = append(errs, "'Path.Notification' is required")
	}
	if t.Form.NotificationPreferences.Frequency == "" {
		errs = append(errs, "'Form.NotificationPreferences.Frequency' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdatePreferenceCommunicationChannelID) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
