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

// GetUploadedMediaFolderForUserCourses Returns the details for a designated upload folder that the user has rights to
// upload to, and creates it if it doesn't exist.
//
// If the current user does not have the permissions to manage files
// in the course or group, the folder will belong to the current user directly.
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type GetUploadedMediaFolderForUserCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetUploadedMediaFolderForUserCourses) GetMethod() string {
	return "GET"
}

func (t *GetUploadedMediaFolderForUserCourses) GetURLPath() string {
	path := "courses/{course_id}/folders/media"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *GetUploadedMediaFolderForUserCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *GetUploadedMediaFolderForUserCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetUploadedMediaFolderForUserCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetUploadedMediaFolderForUserCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetUploadedMediaFolderForUserCourses) Do(c *canvasapi.Canvas) (*models.Folder, error) {
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
