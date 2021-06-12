package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdatePreferencesByCategory Change the preferences for multiple notifications based on the category for a single communication channel
// https://canvas.instructure.com/doc/api/notification_preferences.html
//
// Path Parameters:
// # CommunicationChannelID (Required) ID
// # Category (Required) The name of the category. Must be parameterized (e.g. The category "Course Content" should be "course_content")
//
// Form Parameters:
// # NotificationPreferences (Required) The desired frequency for each notification in the category
//
type UpdatePreferencesByCategory struct {
	Path struct {
		CommunicationChannelID string `json:"communication_channel_id"` //  (Required)
		Category               string `json:"category"`                 //  (Required)
	} `json:"path"`

	Form struct {
		NotificationPreferences struct {
			Frequency string `json:"frequency"` //  (Required)
		} `json:"notification_preferences"`
	} `json:"form"`
}

func (t *UpdatePreferencesByCategory) GetMethod() string {
	return "PUT"
}

func (t *UpdatePreferencesByCategory) GetURLPath() string {
	path := "users/self/communication_channels/{communication_channel_id}/notification_preference_categories/{category}"
	path = strings.ReplaceAll(path, "{communication_channel_id}", fmt.Sprintf("%v", t.Path.CommunicationChannelID))
	path = strings.ReplaceAll(path, "{category}", fmt.Sprintf("%v", t.Path.Category))
	return path
}

func (t *UpdatePreferencesByCategory) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdatePreferencesByCategory) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdatePreferencesByCategory) HasErrors() error {
	errs := []string{}
	if t.Path.CommunicationChannelID == "" {
		errs = append(errs, "'CommunicationChannelID' is required")
	}
	if t.Path.Category == "" {
		errs = append(errs, "'Category' is required")
	}
	if t.Form.NotificationPreferences.Frequency == "" {
		errs = append(errs, "'NotificationPreferences' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdatePreferencesByCategory) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}