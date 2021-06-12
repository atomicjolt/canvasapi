package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ListTopicEntriesCourses Retrieve the (paginated) top-level entries in a discussion topic.
//
// May require (depending on the topic) that the user has posted in the topic.
// If it is required, and the user has not posted, will respond with a 403
// Forbidden status and the body 'require_initial_post'.
//
// Will include the 10 most recent replies, if any, for each entry returned.
//
// If the topic is a root topic with children corresponding to groups of a
// group assignment, entries from those subtopics for which the user belongs
// to the corresponding group will be returned.
//
// Ordering of returned entries is newest-first by posting timestamp (reply
// activity is ignored).
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # CourseID (Required) ID
// # TopicID (Required) ID
//
type ListTopicEntriesCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		TopicID  string `json:"topic_id"`  //  (Required)
	} `json:"path"`
}

func (t *ListTopicEntriesCourses) GetMethod() string {
	return "GET"
}

func (t *ListTopicEntriesCourses) GetURLPath() string {
	path := "courses/{course_id}/discussion_topics/{topic_id}/entries"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	return path
}

func (t *ListTopicEntriesCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ListTopicEntriesCourses) GetBody() (string, error) {
	return "", nil
}

func (t *ListTopicEntriesCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.TopicID == "" {
		errs = append(errs, "'TopicID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListTopicEntriesCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
