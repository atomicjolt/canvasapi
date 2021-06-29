package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// MarkEntryAsUnreadGroups Mark a discussion entry as unread.
//
// No request fields are necessary.
//
// On success, the response will be 204 No Content with an empty body.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # GroupID (Required) ID
// # TopicID (Required) ID
// # EntryID (Required) ID
//
// Query Parameters:
// # ForcedReadState (Optional) A boolean value to set the entry's forced_read_state. No change is made if
//    this argument is not specified.
//
type MarkEntryAsUnreadGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
		TopicID string `json:"topic_id" url:"topic_id,omitempty"` //  (Required)
		EntryID string `json:"entry_id" url:"entry_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		ForcedReadState bool `json:"forced_read_state" url:"forced_read_state,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *MarkEntryAsUnreadGroups) GetMethod() string {
	return "DELETE"
}

func (t *MarkEntryAsUnreadGroups) GetURLPath() string {
	path := "groups/{group_id}/discussion_topics/{topic_id}/entries/{entry_id}/read"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	path = strings.ReplaceAll(path, "{entry_id}", fmt.Sprintf("%v", t.Path.EntryID))
	return path
}

func (t *MarkEntryAsUnreadGroups) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *MarkEntryAsUnreadGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *MarkEntryAsUnreadGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *MarkEntryAsUnreadGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if t.Path.TopicID == "" {
		errs = append(errs, "'TopicID' is required")
	}
	if t.Path.EntryID == "" {
		errs = append(errs, "'EntryID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *MarkEntryAsUnreadGroups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
