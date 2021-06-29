package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ListOfPreferenceCategories Fetch all notification preference categories for the given communication channel
// https://canvas.instructure.com/doc/api/notification_preferences.html
//
// Path Parameters:
// # Path.UserID (Required) ID
// # Path.CommunicationChannelID (Required) ID
//
type ListOfPreferenceCategories struct {
	Path struct {
		UserID                 string `json:"user_id" url:"user_id,omitempty"`                                   //  (Required)
		CommunicationChannelID string `json:"communication_channel_id" url:"communication_channel_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListOfPreferenceCategories) GetMethod() string {
	return "GET"
}

func (t *ListOfPreferenceCategories) GetURLPath() string {
	path := "users/{user_id}/communication_channels/{communication_channel_id}/notification_preference_categories"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{communication_channel_id}", fmt.Sprintf("%v", t.Path.CommunicationChannelID))
	return path
}

func (t *ListOfPreferenceCategories) GetQuery() (string, error) {
	return "", nil
}

func (t *ListOfPreferenceCategories) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListOfPreferenceCategories) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListOfPreferenceCategories) HasErrors() error {
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

func (t *ListOfPreferenceCategories) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
