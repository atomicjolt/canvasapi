package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// ReorderPinnedTopicsCourses Puts the pinned discussion topics in the specified order.
// All pinned topics should be included.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # Order (Required) The ids of the pinned discussion topics in the desired order.
//    (For example, "order=104,102,103".)
//
type ReorderPinnedTopicsCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Order []int64 `json:"order"` //  (Required)
	} `json:"form"`
}

func (t *ReorderPinnedTopicsCourses) GetMethod() string {
	return "POST"
}

func (t *ReorderPinnedTopicsCourses) GetURLPath() string {
	path := "courses/{course_id}/discussion_topics/reorder"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ReorderPinnedTopicsCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ReorderPinnedTopicsCourses) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *ReorderPinnedTopicsCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Form.Order == nil {
		errs = append(errs, "'Order' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ReorderPinnedTopicsCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
