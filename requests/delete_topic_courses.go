package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// DeleteTopicCourses Deletes the discussion topic. This will also delete the assignment, if it's
// an assignment discussion.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # CourseID (Required) ID
// # TopicID (Required) ID
//
type DeleteTopicCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		TopicID  string `json:"topic_id"`  //  (Required)
	} `json:"path"`
}

func (t *DeleteTopicCourses) GetMethod() string {
	return "DELETE"
}

func (t *DeleteTopicCourses) GetURLPath() string {
	path := "courses/{course_id}/discussion_topics/{topic_id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	return path
}

func (t *DeleteTopicCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteTopicCourses) GetBody() (string, error) {
	return "", nil
}

func (t *DeleteTopicCourses) HasErrors() error {
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

func (t *DeleteTopicCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
