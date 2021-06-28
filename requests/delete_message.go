package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// DeleteMessage Delete messages from this conversation. Note that this only affects this
// user's view of the conversation. If all messages are deleted, the
// conversation will be as well (equivalent to DELETE)
// https://canvas.instructure.com/doc/api/conversations.html
//
// Path Parameters:
// # ID (Required) ID
//
// Form Parameters:
// # Remove (Required) Array of message ids to be deleted
//
type DeleteMessage struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Remove []string `json:"remove" url:"remove,omitempty"` //  (Required)
	} `json:"form"`
}

func (t *DeleteMessage) GetMethod() string {
	return "POST"
}

func (t *DeleteMessage) GetURLPath() string {
	path := "conversations/{id}/remove_messages"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteMessage) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteMessage) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *DeleteMessage) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *DeleteMessage) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Form.Remove == nil {
		errs = append(errs, "'Remove' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteMessage) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
