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

// ListPreferencesCommunicationChannelID Fetch all preferences for the given communication channel
// https://canvas.instructure.com/doc/api/notification_preferences.html
//
// Path Parameters:
// # Path.UserID (Required) ID
// # Path.CommunicationChannelID (Required) ID
//
type ListPreferencesCommunicationChannelID struct {
	Path struct {
		UserID                 string `json:"user_id" url:"user_id,omitempty"`                                   //  (Required)
		CommunicationChannelID string `json:"communication_channel_id" url:"communication_channel_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListPreferencesCommunicationChannelID) GetMethod() string {
	return "GET"
}

func (t *ListPreferencesCommunicationChannelID) GetURLPath() string {
	path := "users/{user_id}/communication_channels/{communication_channel_id}/notification_preferences"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{communication_channel_id}", fmt.Sprintf("%v", t.Path.CommunicationChannelID))
	return path
}

func (t *ListPreferencesCommunicationChannelID) GetQuery() (string, error) {
	return "", nil
}

func (t *ListPreferencesCommunicationChannelID) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListPreferencesCommunicationChannelID) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListPreferencesCommunicationChannelID) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if t.Path.CommunicationChannelID == "" {
		errs = append(errs, "'Path.CommunicationChannelID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListPreferencesCommunicationChannelID) Do(c *canvasapi.Canvas) ([]*models.NotificationPreference, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.NotificationPreference{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
