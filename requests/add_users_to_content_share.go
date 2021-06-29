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
)

// AddUsersToContentShare Send a previously created content share to additional users
// https://canvas.instructure.com/doc/api/content_shares.html
//
// Path Parameters:
// # Path.UserID (Required) ID
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.ReceiverIDs (Optional) IDs of users to share the content with.
//
type AddUsersToContentShare struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
		ID     string `json:"id" url:"id,omitempty"`           //  (Required)
	} `json:"path"`

	Form struct {
		ReceiverIDs string `json:"receiver_ids" url:"receiver_ids,omitempty"` //  (Optional)
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

func (t *AddUsersToContentShare) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *AddUsersToContentShare) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *AddUsersToContentShare) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
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
