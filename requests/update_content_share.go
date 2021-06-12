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

// UpdateContentShare Mark a content share read or unread
// https://canvas.instructure.com/doc/api/content_shares.html
//
// Path Parameters:
// # UserID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # ReadState (Optional) . Must be one of read, unreadRead state for the content share
//
type UpdateContentShare struct {
	Path struct {
		UserID string `json:"user_id"` //  (Required)
		ID     string `json:"id"`      //  (Required)
	} `json:"path"`

	Form struct {
		ReadState string `json:"read_state"` //  (Optional) . Must be one of read, unread
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

func (t *UpdateContentShare) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateContentShare) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if !string_utils.Include([]string{"read", "unread"}, t.Form.ReadState) {
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
