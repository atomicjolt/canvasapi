package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ExportGroupsInAndUsersInCategory Returns a csv file of users in format ready to import.
// https://canvas.instructure.com/doc/api/group_categories.html
//
// Path Parameters:
// # GroupCategoryID (Required) ID
//
type ExportGroupsInAndUsersInCategory struct {
	Path struct {
		GroupCategoryID string `json:"group_category_id" url:"group_category_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ExportGroupsInAndUsersInCategory) GetMethod() string {
	return "GET"
}

func (t *ExportGroupsInAndUsersInCategory) GetURLPath() string {
	path := "group_categories/{group_category_id}/export"
	path = strings.ReplaceAll(path, "{group_category_id}", fmt.Sprintf("%v", t.Path.GroupCategoryID))
	return path
}

func (t *ExportGroupsInAndUsersInCategory) GetQuery() (string, error) {
	return "", nil
}

func (t *ExportGroupsInAndUsersInCategory) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ExportGroupsInAndUsersInCategory) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ExportGroupsInAndUsersInCategory) HasErrors() error {
	errs := []string{}
	if t.Path.GroupCategoryID == "" {
		errs = append(errs, "'GroupCategoryID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ExportGroupsInAndUsersInCategory) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
