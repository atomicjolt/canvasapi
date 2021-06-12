package requests

import (
	"fmt"
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
		ID string `json:"id"` //  (Required)
	} `json:"path"`

	Form struct {
		Conversation struct {
			WorkflowState string `json:"workflow_state"` //  (Optional) . Must be one of read, unread, archived
			Subscribed    bool   `json:"subscribed"`     //  (Optional)
			Starred       bool   `json:"starred"`        //  (Optional)
		} `json:"conversation"`

		Scope      string   `json:"scope"`       //  (Optional) . Must be one of unread, starred, archived
		Filter     []string `json:"filter"`      //  (Optional)
		FilterMode string   `json:"filter_mode"` //  (Optional) . Must be one of and, or, default or
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

func (t *EditConversation) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *EditConversation) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if !string_utils.Include([]string{"read", "unread", "archived"}, t.Form.Conversation.WorkflowState) {
		errs = append(errs, "Conversation must be one of read, unread, archived")
	}
	if !string_utils.Include([]string{"unread", "starred", "archived"}, t.Form.Scope) {
		errs = append(errs, "Scope must be one of unread, starred, archived")
	}
	if !string_utils.Include([]string{"and", "or", "default or"}, t.Form.FilterMode) {
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
