package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// UnsubscribeFromTopicCourses Unsubscribe from a topic to stop receiving notifications about new entries
//
// On success, the response will be 204 No Content with an empty body
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # CourseID (Required) ID
// # TopicID (Required) ID
//
type UnsubscribeFromTopicCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		TopicID  string `json:"topic_id"`  //  (Required)
	} `json:"path"`
}

func (t *UnsubscribeFromTopicCourses) GetMethod() string {
	return "DELETE"
}

func (t *UnsubscribeFromTopicCourses) GetURLPath() string {
	path := "courses/{course_id}/discussion_topics/{topic_id}/subscribed"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	return path
}

func (t *UnsubscribeFromTopicCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *UnsubscribeFromTopicCourses) GetBody() (string, error) {
	return "", nil
}

func (t *UnsubscribeFromTopicCourses) HasErrors() error {
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

func (t *UnsubscribeFromTopicCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
