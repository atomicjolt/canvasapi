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

// GetFolderGroups Returns the details for a folder
//
// You can get the root folder from a context by using 'root' as the :id.
// For example, you could get the root folder for a course like:
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # GroupID (Required) ID
// # ID (Required) ID
//
type GetFolderGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
		ID      string `json:"id" url:"id,omitempty"`             //  (Required)
	} `json:"path"`
}

func (t *GetFolderGroups) GetMethod() string {
	return "GET"
}

func (t *GetFolderGroups) GetURLPath() string {
	path := "groups/{group_id}/folders/{id}"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetFolderGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *GetFolderGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetFolderGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetFolderGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetFolderGroups) Do(c *canvasapi.Canvas) (*models.Folder, error) {
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
