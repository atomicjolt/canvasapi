package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
// # UserID (Required) ID
//
// Form Parameters:
// # CommunicationChannel (Required) An email address or SMS number. Not required for "push" type channels.
// # CommunicationChannel (Required) . Must be one of email, sms, pushThe type of communication channel.
//
//    In order to enable push notification support, the server must be
//    properly configured (via sns.yml) to communicate with Amazon
//    Simple Notification Services, and the developer key used to create
//    the access token from this request must have an SNS ARN configured on
//    it.
// # CommunicationChannel (Optional) A registration id, device token, or equivalent token given to an app when
//    registering with a push notification provider. Only valid for "push" type channels.
// # SkipConfirmation (Optional) Only valid for site admins and account admins making requests; If true, the channel is
//    automatically validated and no confirmation email or SMS is sent.
//    Otherwise, the user must respond to a confirmation message to confirm the
//    channel.
//
type CreateCommunicationChannel struct {
	Path struct {
		UserID string `json:"user_id"` //  (Required)
	} `json:"path"`

	Form struct {
		CommunicationChannel struct {
			Address string `json:"address"` //  (Required)
			Type    string `json:"type"`    //  (Required) . Must be one of email, sms, push
			Token   string `json:"token"`   //  (Optional)
		} `json:"communication_channel"`

		SkipConfirmation bool `json:"skip_confirmation"` //  (Optional)
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

func (t *CreateCommunicationChannel) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateCommunicationChannel) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if t.Form.CommunicationChannel.Address == "" {
		errs = append(errs, "'CommunicationChannel' is required")
	}
	if t.Form.CommunicationChannel.Type == "" {
		errs = append(errs, "'CommunicationChannel' is required")
	}
	if !string_utils.Include([]string{"email", "sms", "push"}, t.Form.CommunicationChannel.Type) {
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
