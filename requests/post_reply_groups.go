package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// PostReplyGroups Add a reply to an entry in a discussion topic. Returns a json
// representation of the created reply (see documentation for 'replies'
// method) on success.
//
// May require (depending on the topic) that the user has posted in the topic.
// If it is required, and the user has not posted, will respond with a 403
// Forbidden status and the body 'require_initial_post'.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # GroupID (Required) ID
// # TopicID (Required) ID
// # EntryID (Required) ID
//
// Form Parameters:
// # Message (Optional) The body of the entry.
// # Attachment (Optional) a multipart/form-data form-field-style
//    attachment. Attachments larger than 1 kilobyte are subject to quota
//    restrictions.
//
type PostReplyGroups struct {
	Path struct {
		GroupID string `json:"group_id"` //  (Required)
		TopicID string `json:"topic_id"` //  (Required)
		EntryID string `json:"entry_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Message    string `json:"message"`    //  (Optional)
		Attachment string `json:"attachment"` //  (Optional)
	} `json:"form"`
}

func (t *PostReplyGroups) GetMethod() string {
	return "POST"
}

func (t *PostReplyGroups) GetURLPath() string {
	path := "groups/{group_id}/discussion_topics/{topic_id}/entries/{entry_id}/replies"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	path = strings.ReplaceAll(path, "{entry_id}", fmt.Sprintf("%v", t.Path.EntryID))
	return path
}

func (t *PostReplyGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *PostReplyGroups) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *PostReplyGroups) HasErrors() error {
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

func (t *PostReplyGroups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
