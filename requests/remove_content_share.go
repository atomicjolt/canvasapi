package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// RemoveContentShare Remove a content share from your list. Use +self+ as the user_id. Note that this endpoint does not delete other users'
// copies of the content share.
// https://canvas.instructure.com/doc/api/content_shares.html
//
// Path Parameters:
// # UserID (Required) ID
// # ID (Required) ID
//
type RemoveContentShare struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
		ID     string `json:"id" url:"id,omitempty"`           //  (Required)
	} `json:"path"`
}

func (t *RemoveContentShare) GetMethod() string {
	return "DELETE"
}

func (t *RemoveContentShare) GetURLPath() string {
	path := "users/{user_id}/content_shares/{id}"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *RemoveContentShare) GetQuery() (string, error) {
	return "", nil
}

func (t *RemoveContentShare) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *RemoveContentShare) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *RemoveContentShare) HasErrors() error {
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

func (t *RemoveContentShare) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
