package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetSingleGroupCategory Returns the data for a single group category, or a 401 if the caller doesn't have
// the rights to see it.
// https://canvas.instructure.com/doc/api/group_categories.html
//
// Path Parameters:
// # GroupCategoryID (Required) ID
//
type GetSingleGroupCategory struct {
	Path struct {
		GroupCategoryID string `json:"group_category_id"` //  (Required)
	} `json:"path"`
}

func (t *GetSingleGroupCategory) GetMethod() string {
	return "GET"
}

func (t *GetSingleGroupCategory) GetURLPath() string {
	path := "group_categories/{group_category_id}"
	path = strings.ReplaceAll(path, "{group_category_id}", fmt.Sprintf("%v", t.Path.GroupCategoryID))
	return path
}

func (t *GetSingleGroupCategory) GetQuery() (string, error) {
	return "", nil
}

func (t *GetSingleGroupCategory) GetBody() (string, error) {
	return "", nil
}

func (t *GetSingleGroupCategory) HasErrors() error {
	errs := []string{}
	if t.Path.GroupCategoryID == "" {
		errs = append(errs, "'GroupCategoryID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleGroupCategory) Do(c *canvasapi.Canvas) (*models.GroupCategory, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.GroupCategory{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
