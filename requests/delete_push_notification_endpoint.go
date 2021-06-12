package requests

import (
	"encoding/json"
	"io/ioutil"

	"github.com/atomicjolt/canvasapi"
)

// DeletePushNotificationEndpoint
// https://canvas.instructure.com/doc/api/communication_channels.html
//
type DeletePushNotificationEndpoint struct {
}

func (t *DeletePushNotificationEndpoint) GetMethod() string {
	return "DELETE"
}

func (t *DeletePushNotificationEndpoint) GetURLPath() string {
	return ""
}

func (t *DeletePushNotificationEndpoint) GetQuery() (string, error) {
	return "", nil
}

func (t *DeletePushNotificationEndpoint) GetBody() (string, error) {
	return "", nil
}

func (t *DeletePushNotificationEndpoint) HasErrors() error {
	return nil
}

func (t *DeletePushNotificationEndpoint) Do(c *canvasapi.Canvas) (*canvasapi.SuccessResponse, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := canvasapi.SuccessResponse{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
