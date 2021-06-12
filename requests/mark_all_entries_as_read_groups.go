package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// MarkAllEntriesAsReadGroups Mark the discussion topic and all its entries as read.
//
// No request fields are necessary.
//
// On success, the response will be 204 No Content with an empty body.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # GroupID (Required) ID
// # TopicID (Required) ID
//
// Form Parameters:
// # ForcedReadState (Optional) A boolean value to set all of the entries' forced_read_state. No change
//    is made if this argument is not specified.
//
type MarkAllEntriesAsReadGroups struct {
	Path struct {
		GroupID string `json:"group_id"` //  (Required)
		TopicID string `json:"topic_id"` //  (Required)
	} `json:"path"`

	Form struct {
		ForcedReadState bool `json:"forced_read_state"` //  (Optional)
	} `json:"form"`
}

func (t *MarkAllEntriesAsReadGroups) GetMethod() string {
	return "PUT"
}

func (t *MarkAllEntriesAsReadGroups) GetURLPath() string {
	path := "groups/{group_id}/discussion_topics/{topic_id}/read_all"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	return path
}

func (t *MarkAllEntriesAsReadGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *MarkAllEntriesAsReadGroups) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *MarkAllEntriesAsReadGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if t.Path.TopicID == "" {
		errs = append(errs, "'TopicID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *MarkAllEntriesAsReadGroups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
