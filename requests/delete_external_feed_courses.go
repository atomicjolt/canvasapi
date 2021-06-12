package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// DeleteExternalFeedCourses Deletes the external feed.
// https://canvas.instructure.com/doc/api/announcement_external_feeds.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ExternalFeedID (Required) ID
//
type DeleteExternalFeedCourses struct {
	Path struct {
		CourseID       string `json:"course_id"`        //  (Required)
		ExternalFeedID string `json:"external_feed_id"` //  (Required)
	} `json:"path"`
}

func (t *DeleteExternalFeedCourses) GetMethod() string {
	return "DELETE"
}

func (t *DeleteExternalFeedCourses) GetURLPath() string {
	path := "courses/{course_id}/external_feeds/{external_feed_id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{external_feed_id}", fmt.Sprintf("%v", t.Path.ExternalFeedID))
	return path
}

func (t *DeleteExternalFeedCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteExternalFeedCourses) GetBody() (string, error) {
	return "", nil
}

func (t *DeleteExternalFeedCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ExternalFeedID == "" {
		errs = append(errs, "'ExternalFeedID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteExternalFeedCourses) Do(c *canvasapi.Canvas) (*models.ExternalFeed, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.ExternalFeed{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
