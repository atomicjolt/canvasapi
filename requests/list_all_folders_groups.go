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

// ListAllFoldersGroups Returns the paginated list of all folders for the given context. This will
// be returned as a flat list containing all subfolders as well.
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
//
type ListAllFoldersGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListAllFoldersGroups) GetMethod() string {
	return "GET"
}

func (t *ListAllFoldersGroups) GetURLPath() string {
	path := "groups/{group_id}/folders"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *ListAllFoldersGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *ListAllFoldersGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListAllFoldersGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListAllFoldersGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListAllFoldersGroups) Do(c *canvasapi.Canvas) ([]*models.Folder, error) {
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
