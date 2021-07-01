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

// ListContentExportsUsers A paginated list of the past and pending content export jobs for a course,
// group, or user. Exports are returned newest first.
// https://canvas.instructure.com/doc/api/content_exports.html
//
// Path Parameters:
// # Path.UserID (Required) ID
//
type ListContentExportsUsers struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListContentExportsUsers) GetMethod() string {
	return "GET"
}

func (t *ListContentExportsUsers) GetURLPath() string {
	path := "users/{user_id}/content_exports"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *ListContentExportsUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *ListContentExportsUsers) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListContentExportsUsers) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListContentExportsUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListContentExportsUsers) Do(c *canvasapi.Canvas) ([]*models.ContentExport, *canvasapi.PagedResource, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.ContentExport{}
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
