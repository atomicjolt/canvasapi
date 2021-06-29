package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ListTopicEntriesGroups Retrieve the (paginated) top-level entries in a discussion topic.
//
// May require (depending on the topic) that the user has posted in the topic.
// If it is required, and the user has not posted, will respond with a 403
// Forbidden status and the body 'require_initial_post'.
//
// Will include the 10 most recent replies, if any, for each entry returned.
//
// If the topic is a root topic with children corresponding to groups of a
// group assignment, entries from those subtopics for which the user belongs
// to the corresponding group will be returned.
//
// Ordering of returned entries is newest-first by posting timestamp (reply
// activity is ignored).
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
// # Path.TopicID (Required) ID
//
type ListTopicEntriesGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
		TopicID string `json:"topic_id" url:"topic_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListTopicEntriesGroups) GetMethod() string {
	return "GET"
}

func (t *ListTopicEntriesGroups) GetURLPath() string {
	path := "groups/{group_id}/discussion_topics/{topic_id}/entries"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	return path
}

func (t *ListTopicEntriesGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *ListTopicEntriesGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListTopicEntriesGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListTopicEntriesGroups) HasErrors() error {
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

func (t *ListTopicEntriesGroups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
