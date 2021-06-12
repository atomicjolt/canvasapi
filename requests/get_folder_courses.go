package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetFolderCourses Returns the details for a folder
//
// You can get the root folder from a context by using 'root' as the :id.
// For example, you could get the root folder for a course like:
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
type GetFolderCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`
}

func (t *GetFolderCourses) GetMethod() string {
	return "GET"
}

func (t *GetFolderCourses) GetURLPath() string {
	path := "courses/{course_id}/folders/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetFolderCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *GetFolderCourses) GetBody() (string, error) {
	return "", nil
}

func (t *GetFolderCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetFolderCourses) Do(c *canvasapi.Canvas) (*models.Folder, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Folder{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
