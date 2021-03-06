package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// DeleteEntryCourses Delete a discussion entry.
//
// The entry must have been created by the current user, or the current user
// must have admin rights to the discussion. If the delete is not allowed, a 401 will be returned.
//
// The discussion will be marked deleted, and the user_id and message will be cleared out.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.TopicID (Required) ID
// # Path.ID (Required) ID
//
type DeleteEntryCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		TopicID  string `json:"topic_id" url:"topic_id,omitempty"`   //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`
}

func (t *DeleteEntryCourses) GetMethod() string {
	return "DELETE"
}

func (t *DeleteEntryCourses) GetURLPath() string {
	path := "courses/{course_id}/discussion_topics/{topic_id}/entries/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteEntryCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteEntryCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteEntryCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteEntryCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
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

func (t *DeleteEntryCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
