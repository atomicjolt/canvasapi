package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListAvatarOptions A paginated list of the possible user avatar options that can be set with the user update endpoint. The response will be an array of avatar records. If the 'type' field is 'attachment', the record will include all the normal attachment json fields; otherwise it will include only the 'url' and 'display_name' fields. Additionally, all records will include a 'type' field and a 'token' field. The following explains each field in more detail
// type:: ["gravatar"|"attachment"|"no_pic"] The type of avatar record, for categorization purposes.
// url:: The url of the avatar
// token:: A unique representation of the avatar record which can be used to set the avatar with the user update endpoint. Note: this is an internal representation and is subject to change without notice. It should be consumed with this api endpoint and used in the user update endpoint, and should not be constructed by the client.
// display_name:: A textual description of the avatar record
// id:: ['attachment' type only] the internal id of the attachment
// content-type:: ['attachment' type only] the content-type of the attachment
// filename:: ['attachment' type only] the filename of the attachment
// size:: ['attachment' type only] the size of the attachment
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # UserID (Required) ID
//
type ListAvatarOptions struct {
	Path struct {
		UserID string `json:"user_id"` //  (Required)
	} `json:"path"`
}

func (t *ListAvatarOptions) GetMethod() string {
	return "GET"
}

func (t *ListAvatarOptions) GetURLPath() string {
	path := "users/{user_id}/avatars"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *ListAvatarOptions) GetQuery() (string, error) {
	return "", nil
}

func (t *ListAvatarOptions) GetBody() (string, error) {
	return "", nil
}

func (t *ListAvatarOptions) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListAvatarOptions) Do(c *canvasapi.Canvas) ([]*models.Avatar, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Avatar{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
