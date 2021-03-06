package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetPreferenceType Fetch the preference for the given notification for the given communication channel
// https://canvas.instructure.com/doc/api/notification_preferences.html
//
// Path Parameters:
// # Path.UserID (Required) ID
// # Path.Type (Required) ID
// # Path.Address (Required) ID
// # Path.Notification (Required) ID
//
type GetPreferenceType struct {
	Path struct {
		UserID       string `json:"user_id" url:"user_id,omitempty"`           //  (Required)
		Type         string `json:"type" url:"type,omitempty"`                 //  (Required)
		Address      string `json:"address" url:"address,omitempty"`           //  (Required)
		Notification string `json:"notification" url:"notification,omitempty"` //  (Required)
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

func (t *GetPreferenceType) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetPreferenceType) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetPreferenceType) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if t.Path.Type == "" {
		errs = append(errs, "'Path.Type' is required")
	}
	if t.Path.Address == "" {
		errs = append(errs, "'Path.Address' is required")
	}
	if t.Path.Notification == "" {
		errs = append(errs, "'Path.Notification' is required")
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
