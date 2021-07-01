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

// ResolvePathUsersFullPath Given the full path to a folder, returns a list of all Folders in the path hierarchy,
// starting at the root folder, and ending at the requested folder. The given path is
// relative to the context's root folder and does not include the root folder's name
// (e.g., "course files"). If an empty path is given, the context's root folder alone
// is returned. Otherwise, if no folder exists with the given full path, a Not Found
// error is returned.
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # Path.UserID (Required) ID
//
type ResolvePathUsersFullPath struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ResolvePathUsersFullPath) GetMethod() string {
	return "GET"
}

func (t *ResolvePathUsersFullPath) GetURLPath() string {
	path := "users/{user_id}/folders/by_path/*full_path"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *ResolvePathUsersFullPath) GetQuery() (string, error) {
	return "", nil
}

func (t *ResolvePathUsersFullPath) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ResolvePathUsersFullPath) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ResolvePathUsersFullPath) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ResolvePathUsersFullPath) Do(c *canvasapi.Canvas) ([]*models.Folder, *canvasapi.PagedResource, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.Folder{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, nil, err
	}

	pagedResource, err := canvasapi.ExtractPagedResource(response.Header)
	if err != nil {
		return nil, nil, err
	}

	return ret, pagedResource, nil
}
