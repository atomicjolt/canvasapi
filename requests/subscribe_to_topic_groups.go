package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// SubscribeToTopicGroups Subscribe to a topic to receive notifications about new entries
//
// On success, the response will be 204 No Content with an empty body
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # GroupID (Required) ID
// # TopicID (Required) ID
//
type SubscribeToTopicGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
		TopicID string `json:"topic_id" url:"topic_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *SubscribeToTopicGroups) GetMethod() string {
	return "PUT"
}

func (t *SubscribeToTopicGroups) GetURLPath() string {
	path := "groups/{group_id}/discussion_topics/{topic_id}/subscribed"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	return path
}

func (t *SubscribeToTopicGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *SubscribeToTopicGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *SubscribeToTopicGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *SubscribeToTopicGroups) HasErrors() error {
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

func (t *SubscribeToTopicGroups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
