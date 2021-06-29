package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// CoursesPreviewProcessedHtml Preview html content processed for this course
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
// Form Parameters:
// # Form.Html (Optional) The html content to process
//
type CoursesPreviewProcessedHtml struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Html string `json:"html" url:"html,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *CoursesPreviewProcessedHtml) GetMethod() string {
	return "POST"
}

func (t *CoursesPreviewProcessedHtml) GetURLPath() string {
	path := "courses/{course_id}/preview_html"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CoursesPreviewProcessedHtml) GetQuery() (string, error) {
	return "", nil
}

func (t *CoursesPreviewProcessedHtml) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CoursesPreviewProcessedHtml) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CoursesPreviewProcessedHtml) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CoursesPreviewProcessedHtml) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
