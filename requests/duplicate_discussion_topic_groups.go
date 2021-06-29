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

// DuplicateDiscussionTopicGroups Duplicate a discussion topic according to context (Course/Group)
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
// # Path.TopicID (Required) ID
//
type DuplicateDiscussionTopicGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
		TopicID string `json:"topic_id" url:"topic_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *DuplicateDiscussionTopicGroups) GetMethod() string {
	return "POST"
}

func (t *DuplicateDiscussionTopicGroups) GetURLPath() string {
	path := "groups/{group_id}/discussion_topics/{topic_id}/duplicate"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	return path
}

func (t *DuplicateDiscussionTopicGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *DuplicateDiscussionTopicGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DuplicateDiscussionTopicGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DuplicateDiscussionTopicGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if t.Path.TopicID == "" {
		errs = append(errs, "'Path.TopicID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DuplicateDiscussionTopicGroups) Do(c *canvasapi.Canvas) (*models.DiscussionTopic, error) {
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
