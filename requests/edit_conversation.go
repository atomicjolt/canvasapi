package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// EditConversation Updates attributes for a single conversation.
// https://canvas.instructure.com/doc/api/conversations.html
//
// Path Parameters:
// # ID (Required) ID
//
// Form Parameters:
// # Conversation (Optional) . Must be one of read, unread, archivedChange the state of this conversation
// # Conversation (Optional) Toggle the current user's subscription to the conversation (only valid for
//    group conversations). If unsubscribed, the user will still have access to
//    the latest messages, but the conversation won't be automatically flagged
//    as unread, nor will it jump to the top of the inbox.
// # Conversation (Optional) Toggle the starred state of the current user's view of the conversation.
// # Scope (Optional) . Must be one of unread, starred, archivedUsed when generating "visible" in the API response. See the explanation
//    under the {api:ConversationsController#index index API action}
// # Filter (Optional) Used when generating "visible" in the API response. See the explanation
//    under the {api:ConversationsController#index index API action}
// # FilterMode (Optional) . Must be one of and, or, default orUsed when generating "visible" in the API response. See the explanation
//    under the {api:ConversationsController#index index API action}
//
type EditConversation struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Conversation struct {
			WorkflowState string `json:"workflow_state" url:"workflow_state,omitempty"` //  (Optional) . Must be one of read, unread, archived
			Subscribed    bool   `json:"subscribed" url:"subscribed,omitempty"`         //  (Optional)
			Starred       bool   `json:"starred" url:"starred,omitempty"`               //  (Optional)
		} `json:"conversation" url:"conversation,omitempty"`

		Scope      string   `json:"scope" url:"scope,omitempty"`             //  (Optional) . Must be one of unread, starred, archived
		Filter     []string `json:"filter" url:"filter,omitempty"`           //  (Optional)
		FilterMode string   `json:"filter_mode" url:"filter_mode,omitempty"` //  (Optional) . Must be one of and, or, default or
	} `json:"form"`
}

func (t *EditConversation) GetMethod() string {
	return "PUT"
}

func (t *EditConversation) GetURLPath() string {
	path := "conversations/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *EditConversation) GetQuery() (string, error) {
	return "", nil
}

func (t *EditConversation) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *EditConversation) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *EditConversation) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Form.Conversation.WorkflowState != "" && !string_utils.Include([]string{"read", "unread", "archived"}, t.Form.Conversation.WorkflowState) {
		errs = append(errs, "Conversation must be one of read, unread, archived")
	}
	if t.Form.Scope != "" && !string_utils.Include([]string{"unread", "starred", "archived"}, t.Form.Scope) {
		errs = append(errs, "Scope must be one of unread, starred, archived")
	}
	if t.Form.FilterMode != "" && !string_utils.Include([]string{"and", "or", "default or"}, t.Form.FilterMode) {
		errs = append(errs, "FilterMode must be one of and, or, default or")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *EditConversation) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
