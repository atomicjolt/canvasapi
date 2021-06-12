package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ResolvePathCourses Given the full path to a folder, returns a list of all Folders in the path hierarchy,
// starting at the root folder, and ending at the requested folder. The given path is
// relative to the context's root folder and does not include the root folder's name
// (e.g., "course files"). If an empty path is given, the context's root folder alone
// is returned. Otherwise, if no folder exists with the given full path, a Not Found
// error is returned.
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type ResolvePathCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`
}

func (t *ResolvePathCourses) GetMethod() string {
	return "GET"
}

func (t *ResolvePathCourses) GetURLPath() string {
	path := "courses/{course_id}/folders/by_path"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ResolvePathCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ResolvePathCourses) GetBody() (string, error) {
	return "", nil
}

func (t *ResolvePathCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ResolvePathCourses) Do(c *canvasapi.Canvas) ([]*models.Folder, error) {
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