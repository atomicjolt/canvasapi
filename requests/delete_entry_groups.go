package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// DeleteEntryGroups Delete a discussion entry.
//
// The entry must have been created by the current user, or the current user
// must have admin rights to the discussion. If the delete is not allowed, a 401 will be returned.
//
// The discussion will be marked deleted, and the user_id and message will be cleared out.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
// # Path.TopicID (Required) ID
// # Path.ID (Required) ID
//
type DeleteEntryGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
		TopicID string `json:"topic_id" url:"topic_id,omitempty"` //  (Required)
		ID      string `json:"id" url:"id,omitempty"`             //  (Required)
	} `json:"path"`
}

func (t *DeleteEntryGroups) GetMethod() string {
	return "DELETE"
}

func (t *DeleteEntryGroups) GetURLPath() string {
	path := "groups/{group_id}/discussion_topics/{topic_id}/entries/{id}"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteEntryGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteEntryGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteEntryGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteEntryGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if t.Path.TopicID == "" {
		errs = append(errs, "'Path.TopicID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteEntryGroups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
