package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ShowContentExportGroups Get information about a single content export.
// https://canvas.instructure.com/doc/api/content_exports.html
//
// Path Parameters:
// # GroupID (Required) ID
// # ID (Required) ID
//
type ShowContentExportGroups struct {
	Path struct {
		GroupID string `json:"group_id"` //  (Required)
		ID      string `json:"id"`       //  (Required)
	} `json:"path"`
}

func (t *ShowContentExportGroups) GetMethod() string {
	return "GET"
}

func (t *ShowContentExportGroups) GetURLPath() string {
	path := "groups/{group_id}/content_exports/{id}"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ShowContentExportGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *ShowContentExportGroups) GetBody() (string, error) {
	return "", nil
}

func (t *ShowContentExportGroups) HasErrors() error {
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

func (t *ShowContentExportGroups) Do(c *canvasapi.Canvas) (*models.ContentExport, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.ContentExport{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
