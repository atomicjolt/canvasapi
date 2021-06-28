package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// DeleteConversation Delete this conversation and its messages. Note that this only deletes
// this user's view of the conversation.
//
// Response includes same fields as UPDATE action
// https://canvas.instructure.com/doc/api/conversations.html
//
// Path Parameters:
// # ID (Required) ID
//
type DeleteConversation struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *DeleteConversation) GetMethod() string {
	return "DELETE"
}

func (t *DeleteConversation) GetURLPath() string {
	path := "conversations/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteConversation) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteConversation) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteConversation) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteConversation) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteConversation) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
