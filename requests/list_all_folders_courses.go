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

// ListAllFoldersCourses Returns the paginated list of all folders for the given context. This will
// be returned as a flat list containing all subfolders as well.
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
type ListAllFoldersCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListAllFoldersCourses) GetMethod() string {
	return "GET"
}

func (t *ListAllFoldersCourses) GetURLPath() string {
	path := "courses/{course_id}/folders"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListAllFoldersCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ListAllFoldersCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListAllFoldersCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListAllFoldersCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListAllFoldersCourses) Do(c *canvasapi.Canvas) ([]*models.Folder, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Folder{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
