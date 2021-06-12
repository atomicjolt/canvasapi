package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// DeleteCommunicationChannelID Delete an existing communication channel.
// https://canvas.instructure.com/doc/api/communication_channels.html
//
// Path Parameters:
// # UserID (Required) ID
// # ID (Required) ID
//
type DeleteCommunicationChannelID struct {
	Path struct {
		UserID string `json:"user_id"` //  (Required)
		ID     string `json:"id"`      //  (Required)
	} `json:"path"`
}

func (t *DeleteCommunicationChannelID) GetMethod() string {
	return "DELETE"
}

func (t *DeleteCommunicationChannelID) GetURLPath() string {
	path := "users/{user_id}/communication_channels/{id}"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteCommunicationChannelID) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteCommunicationChannelID) GetBody() (string, error) {
	return "", nil
}

func (t *DeleteCommunicationChannelID) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteCommunicationChannelID) Do(c *canvasapi.Canvas) (*models.CommunicationChannel, error) {
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
