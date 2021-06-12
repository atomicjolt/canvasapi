package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// CoursesPreviewProcessedHtml Preview html content processed for this course
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # Html (Optional) The html content to process
//
type CoursesPreviewProcessedHtml struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Html string `json:"html"` //  (Optional)
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

func (t *CoursesPreviewProcessedHtml) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CoursesPreviewProcessedHtml) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
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
