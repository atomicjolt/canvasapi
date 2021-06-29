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

// DeleteCommunicationChannelType Delete an existing communication channel.
// https://canvas.instructure.com/doc/api/communication_channels.html
//
// Path Parameters:
// # Path.UserID (Required) ID
// # Path.Type (Required) ID
// # Path.Address (Required) ID
//
type DeleteCommunicationChannelType struct {
	Path struct {
		UserID  string `json:"user_id" url:"user_id,omitempty"` //  (Required)
		Type    string `json:"type" url:"type,omitempty"`       //  (Required)
		Address string `json:"address" url:"address,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *DeleteCommunicationChannelType) GetMethod() string {
	return "DELETE"
}

func (t *DeleteCommunicationChannelType) GetURLPath() string {
	path := "users/{user_id}/communication_channels/{type}/{address}"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{type}", fmt.Sprintf("%v", t.Path.Type))
	path = strings.ReplaceAll(path, "{address}", fmt.Sprintf("%v", t.Path.Address))
	return path
}

func (t *DeleteCommunicationChannelType) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteCommunicationChannelType) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteCommunicationChannelType) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteCommunicationChannelType) HasErrors() error {
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
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteCommunicationChannelType) Do(c *canvasapi.Canvas) (*models.CommunicationChannel, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.CommunicationChannel{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
