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

// GetPreferenceCommunicationChannelID Fetch the preference for the given notification for the given communication channel
// https://canvas.instructure.com/doc/api/notification_preferences.html
//
// Path Parameters:
// # Path.UserID (Required) ID
// # Path.CommunicationChannelID (Required) ID
// # Path.Notification (Required) ID
//
type GetPreferenceCommunicationChannelID struct {
	Path struct {
		UserID                 string `json:"user_id" url:"user_id,omitempty"`                                   //  (Required)
		CommunicationChannelID string `json:"communication_channel_id" url:"communication_channel_id,omitempty"` //  (Required)
		Notification           string `json:"notification" url:"notification,omitempty"`                         //  (Required)
	} `json:"path"`
}

func (t *GetPreferenceCommunicationChannelID) GetMethod() string {
	return "GET"
}

func (t *GetPreferenceCommunicationChannelID) GetURLPath() string {
	path := "users/{user_id}/communication_channels/{communication_channel_id}/notification_preferences/{notification}"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{communication_channel_id}", fmt.Sprintf("%v", t.Path.CommunicationChannelID))
	path = strings.ReplaceAll(path, "{notification}", fmt.Sprintf("%v", t.Path.Notification))
	return path
}

func (t *GetPreferenceCommunicationChannelID) GetQuery() (string, error) {
	return "", nil
}

func (t *GetPreferenceCommunicationChannelID) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetPreferenceCommunicationChannelID) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetPreferenceCommunicationChannelID) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if t.Path.CommunicationChannelID == "" {
		errs = append(errs, "'Path.CommunicationChannelID' is required")
	}
	if t.Path.Notification == "" {
		errs = append(errs, "'Path.Notification' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetPreferenceCommunicationChannelID) Do(c *canvasapi.Canvas) (*models.NotificationPreference, error) {
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
