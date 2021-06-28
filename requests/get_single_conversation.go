package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// GetSingleConversation Returns information for a single conversation for the current user. Response includes all
// fields that are present in the list/index action as well as messages
// and extended participant information.
// https://canvas.instructure.com/doc/api/conversations.html
//
// Path Parameters:
// # ID (Required) ID
//
// Query Parameters:
// # InterleaveSubmissions (Optional) (Obsolete) Submissions are no
//    longer linked to conversations. This parameter is ignored.
// # Scope (Optional) . Must be one of unread, starred, archivedUsed when generating "visible" in the API response. See the explanation
//    under the {api:ConversationsController#index index API action}
// # Filter (Optional) Used when generating "visible" in the API response. See the explanation
//    under the {api:ConversationsController#index index API action}
// # FilterMode (Optional) . Must be one of and, or, default orUsed when generating "visible" in the API response. See the explanation
//    under the {api:ConversationsController#index index API action}
// # AutoMarkAsRead (Optional) Default true. If true, unread
//    conversations will be automatically marked as read. This will default
//    to false in a future API release, so clients should explicitly send
//    true if that is the desired behavior.
//
type GetSingleConversation struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		InterleaveSubmissions bool     `json:"interleave_submissions" url:"interleave_submissions,omitempty"` //  (Optional)
		Scope                 string   `json:"scope" url:"scope,omitempty"`                                   //  (Optional) . Must be one of unread, starred, archived
		Filter                []string `json:"filter" url:"filter,omitempty"`                                 //  (Optional)
		FilterMode            string   `json:"filter_mode" url:"filter_mode,omitempty"`                       //  (Optional) . Must be one of and, or, default or
		AutoMarkAsRead        bool     `json:"auto_mark_as_read" url:"auto_mark_as_read,omitempty"`           //  (Optional)
	} `json:"query"`
}

func (t *GetSingleConversation) GetMethod() string {
	return "GET"
}

func (t *GetSingleConversation) GetURLPath() string {
	path := "conversations/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetSingleConversation) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetSingleConversation) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSingleConversation) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSingleConversation) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Query.Scope != "" && !string_utils.Include([]string{"unread", "starred", "archived"}, t.Query.Scope) {
		errs = append(errs, "Scope must be one of unread, starred, archived")
	}
	if t.Query.FilterMode != "" && !string_utils.Include([]string{"and", "or", "default or"}, t.Query.FilterMode) {
		errs = append(errs, "FilterMode must be one of and, or, default or")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleConversation) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
