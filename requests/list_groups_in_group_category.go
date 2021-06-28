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

// ListGroupsInGroupCategory Returns a paginated list of groups in a group category
// https://canvas.instructure.com/doc/api/group_categories.html
//
// Path Parameters:
// # GroupCategoryID (Required) ID
//
type ListGroupsInGroupCategory struct {
	Path struct {
		GroupCategoryID string `json:"group_category_id" url:"group_category_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListGroupsInGroupCategory) GetMethod() string {
	return "GET"
}

func (t *ListGroupsInGroupCategory) GetURLPath() string {
	path := "group_categories/{group_category_id}/groups"
	path = strings.ReplaceAll(path, "{group_category_id}", fmt.Sprintf("%v", t.Path.GroupCategoryID))
	return path
}

func (t *ListGroupsInGroupCategory) GetQuery() (string, error) {
	return "", nil
}

func (t *ListGroupsInGroupCategory) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListGroupsInGroupCategory) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListGroupsInGroupCategory) HasErrors() error {
	errs := []string{}
	if t.Path.GroupCategoryID == "" {
		errs = append(errs, "'GroupCategoryID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListGroupsInGroupCategory) Do(c *canvasapi.Canvas) ([]*models.Group, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Group{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
