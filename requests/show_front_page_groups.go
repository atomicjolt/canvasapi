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

// ShowFrontPageGroups Retrieve the content of the front page
// https://canvas.instructure.com/doc/api/pages.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
//
type ShowFrontPageGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ShowFrontPageGroups) GetMethod() string {
	return "GET"
}

func (t *ShowFrontPageGroups) GetURLPath() string {
	path := "groups/{group_id}/front_page"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *ShowFrontPageGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *ShowFrontPageGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ShowFrontPageGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ShowFrontPageGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowFrontPageGroups) Do(c *canvasapi.Canvas) (*models.Page, error) {
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
