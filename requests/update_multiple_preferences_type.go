package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdateMultiplePreferencesType Change the preferences for multiple notifications for a single communication channel at once
// https://canvas.instructure.com/doc/api/notification_preferences.html
//
// Path Parameters:
// # Type (Required) ID
// # Address (Required) ID
//
// Form Parameters:
// # NotificationPreferences (Required) The desired frequency for <X> notification
//
type UpdateMultiplePreferencesType struct {
	Path struct {
		Type    string `json:"type"`    //  (Required)
		Address string `json:"address"` //  (Required)
	} `json:"path"`

	Form struct {
		NotificationPreferences map[string]UpdateMultiplePreferencesTypeNotificationPreferences
	} `json:"form"`
}

func (t *UpdateMultiplePreferencesType) GetMethod() string {
	return "PUT"
}

func (t *UpdateMultiplePreferencesType) GetURLPath() string {
	path := "users/self/communication_channels/{type}/{address}/notification_preferences"
	path = strings.ReplaceAll(path, "{type}", fmt.Sprintf("%v", t.Path.Type))
	path = strings.ReplaceAll(path, "{address}", fmt.Sprintf("%v", t.Path.Address))
	return path
}

func (t *UpdateMultiplePreferencesType) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateMultiplePreferencesType) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateMultiplePreferencesType) HasErrors() error {
	errs := []string{}
	if t.Path.Type == "" {
		errs = append(errs, "'Type' is required")
	}
	if t.Path.Address == "" {
		errs = append(errs, "'Address' is required")
	}
	if t.Form.NotificationPreferences == nil {
		errs = append(errs, "'NotificationPreferences' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateMultiplePreferencesType) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}

type UpdateMultiplePreferencesTypeNotificationPreferences struct {
	Frequency string `json:"frequency"` //  (Required)
}
