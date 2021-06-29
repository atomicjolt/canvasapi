package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// PostReplyCourses Add a reply to an entry in a discussion topic. Returns a json
// representation of the created reply (see documentation for 'replies'
// method) on success.
//
// May require (depending on the topic) that the user has posted in the topic.
// If it is required, and the user has not posted, will respond with a 403
// Forbidden status and the body 'require_initial_post'.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.TopicID (Required) ID
// # Path.EntryID (Required) ID
//
// Form Parameters:
// # Form.Message (Optional) The body of the entry.
// # Form.Attachment (Optional) a multipart/form-data form-field-style
//    attachment. Attachments larger than 1 kilobyte are subject to quota
//    restrictions.
//
type PostReplyCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		TopicID  string `json:"topic_id" url:"topic_id,omitempty"`   //  (Required)
		EntryID  string `json:"entry_id" url:"entry_id,omitempty"`   //  (Required)
	} `json:"path"`

	Form struct {
		Message    string `json:"message" url:"message,omitempty"`       //  (Optional)
		Attachment string `json:"attachment" url:"attachment,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *PostReplyCourses) GetMethod() string {
	return "POST"
}

func (t *PostReplyCourses) GetURLPath() string {
	path := "courses/{course_id}/discussion_topics/{topic_id}/entries/{entry_id}/replies"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	path = strings.ReplaceAll(path, "{entry_id}", fmt.Sprintf("%v", t.Path.EntryID))
	return path
}

func (t *PostReplyCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *PostReplyCourses) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *PostReplyCourses) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *PostReplyCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.TopicID == "" {
		errs = append(errs, "'Path.TopicID' is required")
	}
	if t.Path.EntryID == "" {
		errs = append(errs, "'Path.EntryID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *PostReplyCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
