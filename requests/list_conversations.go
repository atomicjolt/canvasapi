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

// ListConversations Returns the paginated list of conversations for the current user, most
// recent ones first.
// https://canvas.instructure.com/doc/api/conversations.html
//
// Query Parameters:
// # Query.Scope (Optional) . Must be one of unread, starred, archivedWhen set, only return conversations of the specified type. For example,
//    set to "unread" to return only conversations that haven't been read.
//    The default behavior is to return all non-archived conversations (i.e.
//    read and unread).
// # Query.Filter (Optional) When set, only return conversations for the specified courses, groups
//    or users. The id should be prefixed with its type, e.g. "user_123" or
//    "course_456". Can be an array (by setting "filter[]") or single value
//    (by setting "filter")
// # Query.FilterMode (Optional) . Must be one of and, or, default orWhen filter[] contains multiple filters, combine them with this mode,
//    filtering conversations that at have at least all of the contexts ("and")
//    or at least one of the contexts ("or")
// # Query.InterleaveSubmissions (Optional) (Obsolete) Submissions are no
//    longer linked to conversations. This parameter is ignored.
// # Query.IncludeAllConversationIDs (Optional) Default is false. If true,
//    the top-level element of the response will be an object rather than
//    an array, and will have the keys "conversations" which will contain the
//    paged conversation data, and "conversation_ids" which will contain the
//    ids of all conversations under this scope/filter in the same order.
// # Query.Include (Optional) . Must be one of participant_avatars"participant_avatars":: Optionally include an "avatar_url" key for each user participanting in the conversation
//
type ListConversations struct {
	Query struct {
		Scope                     string   `json:"scope" url:"scope,omitempty"`                                               //  (Optional) . Must be one of unread, starred, archived
		Filter                    []string `json:"filter" url:"filter,omitempty"`                                             //  (Optional)
		FilterMode                string   `json:"filter_mode" url:"filter_mode,omitempty"`                                   //  (Optional) . Must be one of and, or, default or
		InterleaveSubmissions     bool     `json:"interleave_submissions" url:"interleave_submissions,omitempty"`             //  (Optional)
		IncludeAllConversationIDs bool     `json:"include_all_conversation_ids" url:"include_all_conversation_ids,omitempty"` //  (Optional)
		Include                   []string `json:"include" url:"include,omitempty"`                                           //  (Optional) . Must be one of participant_avatars
	} `json:"query"`
}

func (t *ListConversations) GetMethod() string {
	return "GET"
}

func (t *ListConversations) GetURLPath() string {
	return ""
}

func (t *ListConversations) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListConversations) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListConversations) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListConversations) HasErrors() error {
	errs := []string{}
	if t.Query.Scope != "" && !string_utils.Include([]string{"unread", "starred", "archived"}, t.Query.Scope) {
		errs = append(errs, "Scope must be one of unread, starred, archived")
	}
	if t.Query.FilterMode != "" && !string_utils.Include([]string{"and", "or", "default or"}, t.Query.FilterMode) {
		errs = append(errs, "FilterMode must be one of and, or, default or")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"participant_avatars"}, v) {
			errs = append(errs, "Include must be one of participant_avatars")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListConversations) Do(c *canvasapi.Canvas) ([]*models.Conversation, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Conversation{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
