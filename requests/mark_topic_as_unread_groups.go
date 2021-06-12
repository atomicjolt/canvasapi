package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// MarkTopicAsUnreadGroups Mark the initial text of the discussion topic as unread.
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
type MarkTopicAsUnreadGroups struct {
	Path struct {
		GroupID string `json:"group_id"` //  (Required)
		TopicID string `json:"topic_id"` //  (Required)
	} `json:"path"`
}

func (t *MarkTopicAsUnreadGroups) GetMethod() string {
	return "DELETE"
}

func (t *MarkTopicAsUnreadGroups) GetURLPath() string {
	path := "groups/{group_id}/discussion_topics/{topic_id}/read"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	return path
}

func (t *MarkTopicAsUnreadGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *MarkTopicAsUnreadGroups) GetBody() (string, error) {
	return "", nil
}

func (t *MarkTopicAsUnreadGroups) HasErrors() error {
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

func (t *MarkTopicAsUnreadGroups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
