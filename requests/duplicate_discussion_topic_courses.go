package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// DuplicateDiscussionTopicCourses Duplicate a discussion topic according to context (Course/Group)
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.TopicID (Required) ID
//
type DuplicateDiscussionTopicCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		TopicID  string `json:"topic_id" url:"topic_id,omitempty"`   //  (Required)
	} `json:"path"`
}

func (t *DuplicateDiscussionTopicCourses) GetMethod() string {
	return "POST"
}

func (t *DuplicateDiscussionTopicCourses) GetURLPath() string {
	path := "courses/{course_id}/discussion_topics/{topic_id}/duplicate"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	return path
}

func (t *DuplicateDiscussionTopicCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *DuplicateDiscussionTopicCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DuplicateDiscussionTopicCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DuplicateDiscussionTopicCourses) HasErrors() error {
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

func (t *DuplicateDiscussionTopicCourses) Do(c *canvasapi.Canvas) (*models.DiscussionTopic, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.DiscussionTopic{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
