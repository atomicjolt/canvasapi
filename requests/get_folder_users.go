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

// GetFolderUsers Returns the details for a folder
//
// You can get the root folder from a context by using 'root' as the :id.
// For example, you could get the root folder for a course like:
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # Path.UserID (Required) ID
// # Path.ID (Required) ID
//
type GetFolderUsers struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
		ID     string `json:"id" url:"id,omitempty"`           //  (Required)
	} `json:"path"`
}

func (t *GetFolderUsers) GetMethod() string {
	return "GET"
}

func (t *GetFolderUsers) GetURLPath() string {
	path := "users/{user_id}/folders/{id}"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetFolderUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *GetFolderUsers) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetFolderUsers) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetFolderUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetFolderUsers) Do(c *canvasapi.Canvas) (*models.Folder, error) {
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
