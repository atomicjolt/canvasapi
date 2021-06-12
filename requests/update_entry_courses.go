package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdateEntryCourses Update an existing discussion entry.
//
// The entry must have been created by the current user, or the current user
// must have admin rights to the discussion. If the edit is not allowed, a 401 will be returned.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # CourseID (Required) ID
// # TopicID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # Message (Optional) The updated body of the entry.
//
type UpdateEntryCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		TopicID  string `json:"topic_id"`  //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`

	Form struct {
		Message string `json:"message"` //  (Optional)
	} `json:"form"`
}

func (t *UpdateEntryCourses) GetMethod() string {
	return "PUT"
}

func (t *UpdateEntryCourses) GetURLPath() string {
	path := "courses/{course_id}/discussion_topics/{topic_id}/entries/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateEntryCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateEntryCourses) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateEntryCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.TopicID == "" {
		errs = append(errs, "'TopicID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateEntryCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}