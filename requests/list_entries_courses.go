package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// ListEntriesCourses Retrieve a paginated list of discussion entries, given a list of ids.
//
// May require (depending on the topic) that the user has posted in the topic.
// If it is required, and the user has not posted, will respond with a 403
// Forbidden status and the body 'require_initial_post'.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.TopicID (Required) ID
//
// Query Parameters:
// # Query.IDs (Optional) A list of entry ids to retrieve. Entries will be returned in id order,
//    smallest id first.
//
type ListEntriesCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		TopicID  string `json:"topic_id" url:"topic_id,omitempty"`   //  (Required)
	} `json:"path"`

	Query struct {
		IDs []string `json:"ids" url:"ids,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ListEntriesCourses) GetMethod() string {
	return "GET"
}

func (t *ListEntriesCourses) GetURLPath() string {
	path := "courses/{course_id}/discussion_topics/{topic_id}/entry_list"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	return path
}

func (t *ListEntriesCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListEntriesCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListEntriesCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListEntriesCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.TopicID == "" {
		errs = append(errs, "'Path.TopicID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListEntriesCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
