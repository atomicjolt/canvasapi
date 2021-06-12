package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdatePreferenceType Change the preference for a single notification for a single communication channel
// https://canvas.instructure.com/doc/api/notification_preferences.html
//
// Path Parameters:
// # Type (Required) ID
// # Address (Required) ID
// # Notification (Required) ID
//
// Form Parameters:
// # NotificationPreferences (Required) The desired frequency for this notification
//
type UpdatePreferenceType struct {
	Path struct {
		Type         string `json:"type"`         //  (Required)
		Address      string `json:"address"`      //  (Required)
		Notification string `json:"notification"` //  (Required)
	} `json:"path"`

	Form struct {
		NotificationPreferences struct {
			Frequency string `json:"frequency"` //  (Required)
		} `json:"notification_preferences"`
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

func (t *UpdatePreferenceType) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdatePreferenceType) HasErrors() error {
	errs := []string{}
	if t.Path.Type == "" {
		errs = append(errs, "'Type' is required")
	}
	if t.Path.Address == "" {
		errs = append(errs, "'Address' is required")
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

func (t *UpdatePreferenceType) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
