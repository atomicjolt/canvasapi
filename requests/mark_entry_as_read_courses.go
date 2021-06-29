package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// MarkEntryAsReadCourses Mark a discussion entry as read.
//
// No request fields are necessary.
//
// On success, the response will be 204 No Content with an empty body.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.TopicID (Required) ID
// # Path.EntryID (Required) ID
//
// Form Parameters:
// # Form.ForcedReadState (Optional) A boolean value to set the entry's forced_read_state. No change is made if
//    this argument is not specified.
//
type MarkEntryAsReadCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		TopicID  string `json:"topic_id" url:"topic_id,omitempty"`   //  (Required)
		EntryID  string `json:"entry_id" url:"entry_id,omitempty"`   //  (Required)
	} `json:"path"`

	Form struct {
		ForcedReadState bool `json:"forced_read_state" url:"forced_read_state,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *MarkEntryAsReadCourses) GetMethod() string {
	return "PUT"
}

func (t *MarkEntryAsReadCourses) GetURLPath() string {
	path := "courses/{course_id}/discussion_topics/{topic_id}/entries/{entry_id}/read"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	path = strings.ReplaceAll(path, "{entry_id}", fmt.Sprintf("%v", t.Path.EntryID))
	return path
}

func (t *MarkEntryAsReadCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *MarkEntryAsReadCourses) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *MarkEntryAsReadCourses) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *MarkEntryAsReadCourses) HasErrors() error {
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

func (t *MarkEntryAsReadCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
