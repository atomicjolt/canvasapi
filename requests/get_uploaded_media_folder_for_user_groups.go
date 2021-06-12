package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetUploadedMediaFolderForUserGroups Returns the details for a designated upload folder that the user has rights to
// upload to, and creates it if it doesn't exist.
//
// If the current user does not have the permissions to manage files
// in the course or group, the folder will belong to the current user directly.
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # GroupID (Required) ID
//
type GetUploadedMediaFolderForUserGroups struct {
	Path struct {
		GroupID string `json:"group_id"` //  (Required)
	} `json:"path"`
}

func (t *GetUploadedMediaFolderForUserGroups) GetMethod() string {
	return "GET"
}

func (t *GetUploadedMediaFolderForUserGroups) GetURLPath() string {
	path := "groups/{group_id}/folders/media"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *GetUploadedMediaFolderForUserGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *GetUploadedMediaFolderForUserGroups) GetBody() (string, error) {
	return "", nil
}

func (t *GetUploadedMediaFolderForUserGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetUploadedMediaFolderForUserGroups) Do(c *canvasapi.Canvas) (*models.Folder, error) {
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
