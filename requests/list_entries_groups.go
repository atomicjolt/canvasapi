package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// ListEntriesGroups Retrieve a paginated list of discussion entries, given a list of ids.
//
// May require (depending on the topic) that the user has posted in the topic.
// If it is required, and the user has not posted, will respond with a 403
// Forbidden status and the body 'require_initial_post'.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # GroupID (Required) ID
// # TopicID (Required) ID
//
// Query Parameters:
// # IDs (Optional) A list of entry ids to retrieve. Entries will be returned in id order,
//    smallest id first.
//
type ListEntriesGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
		TopicID string `json:"topic_id" url:"topic_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		IDs []string `json:"ids" url:"ids,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ListEntriesGroups) GetMethod() string {
	return "GET"
}

func (t *ListEntriesGroups) GetURLPath() string {
	path := "groups/{group_id}/discussion_topics/{topic_id}/entry_list"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	return path
}

func (t *ListEntriesGroups) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListEntriesGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListEntriesGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListEntriesGroups) HasErrors() error {
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

func (t *ListEntriesGroups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
