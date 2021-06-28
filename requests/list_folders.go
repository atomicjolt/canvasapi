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

// ListFolders Returns the paginated list of folders in the folder.
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # ID (Required) ID
//
type ListFolders struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListFolders) GetMethod() string {
	return "GET"
}

func (t *ListFolders) GetURLPath() string {
	path := "folders/{id}/folders"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ListFolders) GetQuery() (string, error) {
	return "", nil
}

func (t *ListFolders) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListFolders) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListFolders) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListFolders) Do(c *canvasapi.Canvas) ([]*models.Folder, error) {
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
