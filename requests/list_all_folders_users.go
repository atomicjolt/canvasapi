package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListAllFoldersUsers Returns the paginated list of all folders for the given context. This will
// be returned as a flat list containing all subfolders as well.
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # Path.UserID (Required) ID
//
type ListAllFoldersUsers struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListAllFoldersUsers) GetMethod() string {
	return "GET"
}

func (t *ListAllFoldersUsers) GetURLPath() string {
	path := "users/{user_id}/folders"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *ListAllFoldersUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *ListAllFoldersUsers) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListAllFoldersUsers) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListAllFoldersUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListAllFoldersUsers) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.Folder, *canvasapi.PagedResource, error) {
	var err error
	var response *http.Response
	if next != nil {
		response, err = c.Send(next, t.GetMethod(), nil)
	} else {
		response, err = c.SendRequest(t)
	}

	if err != nil {
		return nil, nil, err
	}
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
