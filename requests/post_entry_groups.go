package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// PostEntryGroups Create a new entry in a discussion topic. Returns a json representation of
// the created entry (see documentation for 'entries' method) on success.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # GroupID (Required) ID
// # TopicID (Required) ID
//
// Form Parameters:
// # Message (Optional) The body of the entry.
// # Attachment (Optional) a multipart/form-data form-field-style
//    attachment. Attachments larger than 1 kilobyte are subject to quota
//    restrictions.
//
type PostEntryGroups struct {
	Path struct {
		GroupID string `json:"group_id"` //  (Required)
		TopicID string `json:"topic_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Message    string `json:"message"`    //  (Optional)
		Attachment string `json:"attachment"` //  (Optional)
	} `json:"form"`
}

func (t *PostEntryGroups) GetMethod() string {
	return "POST"
}

func (t *PostEntryGroups) GetURLPath() string {
	path := "groups/{group_id}/discussion_topics/{topic_id}/entries"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	return path
}

func (t *PostEntryGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *PostEntryGroups) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *PostEntryGroups) HasErrors() error {
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

func (t *PostEntryGroups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}