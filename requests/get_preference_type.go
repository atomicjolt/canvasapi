package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetPreferenceType Fetch the preference for the given notification for the given communication channel
// https://canvas.instructure.com/doc/api/notification_preferences.html
//
// Path Parameters:
// # UserID (Required) ID
// # Type (Required) ID
// # Address (Required) ID
// # Notification (Required) ID
//
type GetPreferenceType struct {
	Path struct {
		UserID       string `json:"user_id"`      //  (Required)
		Type         string `json:"type"`         //  (Required)
		Address      string `json:"address"`      //  (Required)
		Notification string `json:"notification"` //  (Required)
	} `json:"path"`
}

func (t *GetPreferenceType) GetMethod() string {
	return "GET"
}

func (t *GetPreferenceType) GetURLPath() string {
	path := "users/{user_id}/communication_channels/{type}/{address}/notification_preferences/{notification}"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{type}", fmt.Sprintf("%v", t.Path.Type))
	path = strings.ReplaceAll(path, "{address}", fmt.Sprintf("%v", t.Path.Address))
	path = strings.ReplaceAll(path, "{notification}", fmt.Sprintf("%v", t.Path.Notification))
	return path
}

func (t *GetPreferenceType) GetQuery() (string, error) {
	return "", nil
}

func (t *GetPreferenceType) GetBody() (string, error) {
	return "", nil
}

func (t *GetPreferenceType) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if t.Path.Type == "" {
		errs = append(errs, "'Type' is required")
	}
	if t.Path.Address == "" {
		errs = append(errs, "'Address' is required")
	}
	if t.Path.Notification == "" {
		errs = append(errs, "'Notification' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetPreferenceType) Do(c *canvasapi.Canvas) (*models.NotificationPreference, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.NotificationPreference{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
