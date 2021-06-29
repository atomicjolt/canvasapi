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

// UpdateContentShare Mark a content share read or unread
// https://canvas.instructure.com/doc/api/content_shares.html
//
// Path Parameters:
// # Path.UserID (Required) ID
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.ReadState (Optional) . Must be one of read, unreadRead state for the content share
//
type UpdateContentShare struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
		ID     string `json:"id" url:"id,omitempty"`           //  (Required)
	} `json:"path"`

	Form struct {
		ReadState string `json:"read_state" url:"read_state,omitempty"` //  (Optional) . Must be one of read, unread
	} `json:"form"`
}

func (t *UpdateContentShare) GetMethod() string {
	return "PUT"
}

func (t *UpdateContentShare) GetURLPath() string {
	path := "users/{user_id}/content_shares/{id}"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateContentShare) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateContentShare) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateContentShare) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateContentShare) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if t.Form.ReadState != "" && !string_utils.Include([]string{"read", "unread"}, t.Form.ReadState) {
		errs = append(errs, "ReadState must be one of read, unread")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateContentShare) Do(c *canvasapi.Canvas) (*models.ContentShare, error) {
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
