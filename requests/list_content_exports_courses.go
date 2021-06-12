package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListContentExportsCourses A paginated list of the past and pending content export jobs for a course,
// group, or user. Exports are returned newest first.
// https://canvas.instructure.com/doc/api/content_exports.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type ListContentExportsCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`
}

func (t *ListContentExportsCourses) GetMethod() string {
	return "GET"
}

func (t *ListContentExportsCourses) GetURLPath() string {
	path := "courses/{course_id}/content_exports"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListContentExportsCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ListContentExportsCourses) GetBody() (string, error) {
	return "", nil
}

func (t *ListContentExportsCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListContentExportsCourses) Do(c *canvasapi.Canvas) ([]*models.ContentExport, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.ContentExport{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
