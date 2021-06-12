package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// FindRecipientsSearch Find valid recipients (users, courses and groups) that the current user
// can send messages to. The /api/v1/search/recipients path is the preferred
// endpoint, /api/v1/conversations/find_recipients is deprecated.
//
// Pagination is supported.
// https://canvas.instructure.com/doc/api/search.html
//
// Query Parameters:
// # Search (Optional) Search terms used for matching users/courses/groups (e.g. "bob smith"). If
//    multiple terms are given (separated via whitespace), only results matching
//    all terms will be returned.
// # Context (Optional) Limit the search to a particular course/group (e.g. "course_3" or "group_4").
// # Exclude (Optional) Array of ids to exclude from the search. These may be user ids or
//    course/group ids prefixed with "course_" or "group_" respectively,
//    e.g. exclude[]=1&exclude[]=2&exclude[]=course_3
// # Type (Optional) . Must be one of user, contextLimit the search just to users or contexts (groups/courses).
// # UserID (Optional) Search for a specific user id. This ignores the other above parameters,
//    and will never return more than one result.
// # FromConversationID (Optional) When searching by user_id, only users that could be normally messaged by
//    this user will be returned. This parameter allows you to specify a
//    conversation that will be referenced for a shared context -- if both the
//    current user and the searched user are in the conversation, the user will
//    be returned. This is used to start new side conversations.
// # Permissions (Optional) Array of permission strings to be checked for each matched context (e.g.
//    "send_messages"). This argument determines which permissions may be
//    returned in the response; it won't prevent contexts from being returned if
//    they don't grant the permission(s).
//
type FindRecipientsSearch struct {
	Query struct {
		Search             string   `json:"search"`               //  (Optional)
		Context            string   `json:"context"`              //  (Optional)
		Exclude            []string `json:"exclude"`              //  (Optional)
		Type               string   `json:"type"`                 //  (Optional) . Must be one of user, context
		UserID             int64    `json:"user_id"`              //  (Optional)
		FromConversationID int64    `json:"from_conversation_id"` //  (Optional)
		Permissions        []string `json:"permissions"`          //  (Optional)
	} `json:"query"`
}

func (t *FindRecipientsSearch) GetMethod() string {
	return "GET"
}

func (t *FindRecipientsSearch) GetURLPath() string {
	return ""
}

func (t *FindRecipientsSearch) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *FindRecipientsSearch) GetBody() (string, error) {
	return "", nil
}

func (t *FindRecipientsSearch) HasErrors() error {
	errs := []string{}
	if !string_utils.Include([]string{"user", "context"}, t.Query.Type) {
		errs = append(errs, "Type must be one of user, context")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *FindRecipientsSearch) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
