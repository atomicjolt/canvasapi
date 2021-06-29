package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdatePreferenceType Change the preference for a single notification for a single communication channel
// https://canvas.instructure.com/doc/api/notification_preferences.html
//
// Path Parameters:
// # Path.Type (Required) ID
// # Path.Address (Required) ID
// # Path.Notification (Required) ID
//
// Form Parameters:
// # Form.NotificationPreferences.Frequency (Required) The desired frequency for this notification
//
type UpdatePreferenceType struct {
	Path struct {
		Type         string `json:"type" url:"type,omitempty"`                 //  (Required)
		Address      string `json:"address" url:"address,omitempty"`           //  (Required)
		Notification string `json:"notification" url:"notification,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		NotificationPreferences struct {
			Frequency string `json:"frequency" url:"frequency,omitempty"` //  (Required)
		} `json:"notification_preferences" url:"notification_preferences,omitempty"`
	} `json:"form"`
}

func (t *UpdatePreferenceType) GetMethod() string {
	return "PUT"
}

func (t *UpdatePreferenceType) GetURLPath() string {
	path := "users/self/communication_channels/{type}/{address}/notification_preferences/{notification}"
	path = strings.ReplaceAll(path, "{type}", fmt.Sprintf("%v", t.Path.Type))
	path = strings.ReplaceAll(path, "{address}", fmt.Sprintf("%v", t.Path.Address))
	path = strings.ReplaceAll(path, "{notification}", fmt.Sprintf("%v", t.Path.Notification))
	return path
}

func (t *UpdatePreferenceType) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdatePreferenceType) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdatePreferenceType) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdatePreferenceType) HasErrors() error {
	errs := []string{}
	if t.Path.Type == "" {
		errs = append(errs, "'Path.Type' is required")
	}
	if t.Path.Address == "" {
		errs = append(errs, "'Path.Address' is required")
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

func (t *UpdatePreferenceType) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
