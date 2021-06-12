package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdatePreferenceCommunicationChannelID Change the preference for a single notification for a single communication channel
// https://canvas.instructure.com/doc/api/notification_preferences.html
//
// Path Parameters:
// # CommunicationChannelID (Required) ID
// # Notification (Required) ID
//
// Form Parameters:
// # NotificationPreferences (Required) The desired frequency for this notification
//
type UpdatePreferenceCommunicationChannelID struct {
	Path struct {
		CommunicationChannelID string `json:"communication_channel_id"` //  (Required)
		Notification           string `json:"notification"`             //  (Required)
	} `json:"path"`

	Form struct {
		NotificationPreferences struct {
			Frequency string `json:"frequency"` //  (Required)
		} `json:"notification_preferences"`
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

func (t *UpdatePreferenceCommunicationChannelID) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdatePreferenceCommunicationChannelID) HasErrors() error {
	errs := []string{}
	if t.Path.CommunicationChannelID == "" {
		errs = append(errs, "'CommunicationChannelID' is required")
	}
	if t.Path.Notification == "" {
		errs = append(errs, "'Notification' is required")
	}
	if t.Form.NotificationPreferences.Frequency == "" {
		errs = append(errs, "'NotificationPreferences' is required")
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
