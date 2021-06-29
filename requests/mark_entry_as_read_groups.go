package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// MarkEntryAsReadGroups Mark a discussion entry as read.
//
// No request fields are necessary.
//
// On success, the response will be 204 No Content with an empty body.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
// # Path.TopicID (Required) ID
// # Path.EntryID (Required) ID
//
// Form Parameters:
// # Form.ForcedReadState (Optional) A boolean value to set the entry's forced_read_state. No change is made if
//    this argument is not specified.
//
type MarkEntryAsReadGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
		TopicID string `json:"topic_id" url:"topic_id,omitempty"` //  (Required)
		EntryID string `json:"entry_id" url:"entry_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		ForcedReadState bool `json:"forced_read_state" url:"forced_read_state,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *MarkEntryAsReadGroups) GetMethod() string {
	return "PUT"
}

func (t *MarkEntryAsReadGroups) GetURLPath() string {
	path := "groups/{group_id}/discussion_topics/{topic_id}/entries/{entry_id}/read"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	path = strings.ReplaceAll(path, "{entry_id}", fmt.Sprintf("%v", t.Path.EntryID))
	return path
}

func (t *MarkEntryAsReadGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *MarkEntryAsReadGroups) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *MarkEntryAsReadGroups) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *MarkEntryAsReadGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if t.Path.TopicID == "" {
		errs = append(errs, "'Path.TopicID' is required")
	}
	if t.Path.EntryID == "" {
		errs = append(errs, "'Path.EntryID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *MarkEntryAsReadGroups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
