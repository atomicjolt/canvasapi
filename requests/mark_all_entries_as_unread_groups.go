package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// MarkAllEntriesAsUnreadGroups Mark the discussion topic and all its entries as unread.
//
// No request fields are necessary.
//
// On success, the response will be 204 No Content with an empty body.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
// # Path.TopicID (Required) ID
//
// Query Parameters:
// # Query.ForcedReadState (Optional) A boolean value to set all of the entries' forced_read_state. No change is
//    made if this argument is not specified.
//
type MarkAllEntriesAsUnreadGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
		TopicID string `json:"topic_id" url:"topic_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		ForcedReadState bool `json:"forced_read_state" url:"forced_read_state,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *MarkAllEntriesAsUnreadGroups) GetMethod() string {
	return "DELETE"
}

func (t *MarkAllEntriesAsUnreadGroups) GetURLPath() string {
	path := "groups/{group_id}/discussion_topics/{topic_id}/read_all"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	return path
}

func (t *MarkAllEntriesAsUnreadGroups) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *MarkAllEntriesAsUnreadGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *MarkAllEntriesAsUnreadGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *MarkAllEntriesAsUnreadGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if t.Path.TopicID == "" {
		errs = append(errs, "'Path.TopicID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *MarkAllEntriesAsUnreadGroups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
