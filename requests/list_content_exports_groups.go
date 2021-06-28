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

// ListContentExportsGroups A paginated list of the past and pending content export jobs for a course,
// group, or user. Exports are returned newest first.
// https://canvas.instructure.com/doc/api/content_exports.html
//
// Path Parameters:
// # GroupID (Required) ID
//
type ListContentExportsGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListContentExportsGroups) GetMethod() string {
	return "GET"
}

func (t *ListContentExportsGroups) GetURLPath() string {
	path := "groups/{group_id}/content_exports"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *ListContentExportsGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *ListContentExportsGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListContentExportsGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListContentExportsGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListContentExportsGroups) Do(c *canvasapi.Canvas) ([]*models.ContentExport, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.ContentExport{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
