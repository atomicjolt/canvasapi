package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ListEntryRepliesGroups Retrieve the (paginated) replies to a top-level entry in a discussion
// topic.
//
// May require (depending on the topic) that the user has posted in the topic.
// If it is required, and the user has not posted, will respond with a 403
// Forbidden status and the body 'require_initial_post'.
//
// Ordering of returned entries is newest-first by creation timestamp.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # GroupID (Required) ID
// # TopicID (Required) ID
// # EntryID (Required) ID
//
type ListEntryRepliesGroups struct {
	Path struct {
		GroupID string `json:"group_id"` //  (Required)
		TopicID string `json:"topic_id"` //  (Required)
		EntryID string `json:"entry_id"` //  (Required)
	} `json:"path"`
}

func (t *ListEntryRepliesGroups) GetMethod() string {
	return "GET"
}

func (t *ListEntryRepliesGroups) GetURLPath() string {
	path := "groups/{group_id}/discussion_topics/{topic_id}/entries/{entry_id}/replies"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	path = strings.ReplaceAll(path, "{entry_id}", fmt.Sprintf("%v", t.Path.EntryID))
	return path
}

func (t *ListEntryRepliesGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *ListEntryRepliesGroups) GetBody() (string, error) {
	return "", nil
}

func (t *ListEntryRepliesGroups) HasErrors() error {
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

func (t *ListEntryRepliesGroups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
