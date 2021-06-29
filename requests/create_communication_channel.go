package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// CreateCommunicationChannel Creates a new communication channel for the specified user.
// https://canvas.instructure.com/doc/api/communication_channels.html
//
// Path Parameters:
// # Path.UserID (Required) ID
//
// Form Parameters:
// # Form.CommunicationChannel.Address (Required) An email address or SMS number. Not required for "push" type channels.
// # Form.CommunicationChannel.Type (Required) . Must be one of email, sms, pushThe type of communication channel.
//
//    In order to enable push notification support, the server must be
//    properly configured (via sns.yml) to communicate with Amazon
//    Simple Notification Services, and the developer key used to create
//    the access token from this request must have an SNS ARN configured on
//    it.
// # Form.CommunicationChannel.Token (Optional) A registration id, device token, or equivalent token given to an app when
//    registering with a push notification provider. Only valid for "push" type channels.
// # Form.SkipConfirmation (Optional) Only valid for site admins and account admins making requests; If true, the channel is
//    automatically validated and no confirmation email or SMS is sent.
//    Otherwise, the user must respond to a confirmation message to confirm the
//    channel.
//
type CreateCommunicationChannel struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		CommunicationChannel struct {
			Address string `json:"address" url:"address,omitempty"` //  (Required)
			Type    string `json:"type" url:"type,omitempty"`       //  (Required) . Must be one of email, sms, push
			Token   string `json:"token" url:"token,omitempty"`     //  (Optional)
		} `json:"communication_channel" url:"communication_channel,omitempty"`

		SkipConfirmation bool `json:"skip_confirmation" url:"skip_confirmation,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *CreateCommunicationChannel) GetMethod() string {
	return "POST"
}

func (t *CreateCommunicationChannel) GetURLPath() string {
	path := "users/{user_id}/communication_channels"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *CreateCommunicationChannel) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateCommunicationChannel) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateCommunicationChannel) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateCommunicationChannel) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if t.Form.CommunicationChannel.Address == "" {
		errs = append(errs, "'Form.CommunicationChannel.Address' is required")
	}
	if t.Form.CommunicationChannel.Type == "" {
		errs = append(errs, "'Form.CommunicationChannel.Type' is required")
	}
	if t.Form.CommunicationChannel.Type != "" && !string_utils.Include([]string{"email", "sms", "push"}, t.Form.CommunicationChannel.Type) {
		errs = append(errs, "CommunicationChannel must be one of email, sms, push")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateCommunicationChannel) Do(c *canvasapi.Canvas) (*models.CommunicationChannel, error) {
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
