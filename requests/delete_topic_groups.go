package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// DeleteTopicGroups Deletes the discussion topic. This will also delete the assignment, if it's
// an assignment discussion.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
// # Path.TopicID (Required) ID
//
type DeleteTopicGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
		TopicID string `json:"topic_id" url:"topic_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *DeleteTopicGroups) GetMethod() string {
	return "DELETE"
}

func (t *DeleteTopicGroups) GetURLPath() string {
	path := "groups/{group_id}/discussion_topics/{topic_id}"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	return path
}

func (t *DeleteTopicGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteTopicGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteTopicGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteTopicGroups) HasErrors() error {
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

func (t *DeleteTopicGroups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
