package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// AddUsersToContentShare Send a previously created content share to additional users
// https://canvas.instructure.com/doc/api/content_shares.html
//
// Path Parameters:
// # UserID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # ReceiverIDs (Optional) IDs of users to share the content with.
//
type AddUsersToContentShare struct {
	Path struct {
		UserID string `json:"user_id"` //  (Required)
		ID     string `json:"id"`      //  (Required)
	} `json:"path"`

	Form struct {
		ReceiverIDs string `json:"receiver_ids"` //  (Optional)
	} `json:"form"`
}

func (t *AddUsersToContentShare) GetMethod() string {
	return "POST"
}

func (t *AddUsersToContentShare) GetURLPath() string {
	path := "users/{user_id}/content_shares/{id}/add_users"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *AddUsersToContentShare) GetQuery() (string, error) {
	return "", nil
}

func (t *AddUsersToContentShare) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *AddUsersToContentShare) HasErrors() error {
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

func (t *AddUsersToContentShare) Do(c *canvasapi.Canvas) (*models.ContentShare, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.ContentShare{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
