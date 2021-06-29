package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// RateEntryCourses Rate a discussion entry.
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
// # Form.Rating (Optional) A rating to set on this entry. Only 0 and 1 are accepted.
//
type RateEntryCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		TopicID  string `json:"topic_id" url:"topic_id,omitempty"`   //  (Required)
		EntryID  string `json:"entry_id" url:"entry_id,omitempty"`   //  (Required)
	} `json:"path"`

	Form struct {
		Rating int64 `json:"rating" url:"rating,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *RateEntryCourses) GetMethod() string {
	return "POST"
}

func (t *RateEntryCourses) GetURLPath() string {
	path := "courses/{course_id}/discussion_topics/{topic_id}/entries/{entry_id}/rating"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	path = strings.ReplaceAll(path, "{entry_id}", fmt.Sprintf("%v", t.Path.EntryID))
	return path
}

func (t *RateEntryCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *RateEntryCourses) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *RateEntryCourses) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *RateEntryCourses) HasErrors() error {
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

func (t *RateEntryCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
