package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// DeleteGroupCategory Deletes a group category and all groups under it. Protected group
// categories can not be deleted, i.e. "communities" and "student_organized".
// https://canvas.instructure.com/doc/api/group_categories.html
//
// Path Parameters:
// # Path.GroupCategoryID (Required) ID
//
type DeleteGroupCategory struct {
	Path struct {
		GroupCategoryID string `json:"group_category_id" url:"group_category_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *DeleteGroupCategory) GetMethod() string {
	return "DELETE"
}

func (t *DeleteGroupCategory) GetURLPath() string {
	path := "group_categories/{group_category_id}"
	path = strings.ReplaceAll(path, "{group_category_id}", fmt.Sprintf("%v", t.Path.GroupCategoryID))
	return path
}

func (t *DeleteGroupCategory) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteGroupCategory) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteGroupCategory) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteGroupCategory) HasErrors() error {
	errs := []string{}
	if t.Path.GroupCategoryID == "" {
		errs = append(errs, "'Path.GroupCategoryID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteGroupCategory) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
