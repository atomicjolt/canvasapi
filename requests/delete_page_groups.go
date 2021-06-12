package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// DeletePageGroups Delete a wiki page
// https://canvas.instructure.com/doc/api/pages.html
//
// Path Parameters:
// # GroupID (Required) ID
// # Url (Required) ID
//
type DeletePageGroups struct {
	Path struct {
		GroupID string `json:"group_id"` //  (Required)
		Url     string `json:"url"`      //  (Required)
	} `json:"path"`
}

func (t *DeletePageGroups) GetMethod() string {
	return "DELETE"
}

func (t *DeletePageGroups) GetURLPath() string {
	path := "groups/{group_id}/pages/{url}"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{url}", fmt.Sprintf("%v", t.Path.Url))
	return path
}

func (t *DeletePageGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *DeletePageGroups) GetBody() (string, error) {
	return "", nil
}

func (t *DeletePageGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if t.Path.Url == "" {
		errs = append(errs, "'Url' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeletePageGroups) Do(c *canvasapi.Canvas) (*models.Page, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Page{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
