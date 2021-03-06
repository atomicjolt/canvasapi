package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// GroupsPermissions Returns permission information for the calling user in the given group.
// See also the {api:AccountsController#permissions Account} and
// {api:CoursesController#permissions Course} counterparts.
// https://canvas.instructure.com/doc/api/groups.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
//
// Query Parameters:
// # Query.Permissions (Optional) List of permissions to check against the authenticated user.
//    Permission names are documented in the {api:RoleOverridesController#add_role Create a role} endpoint.
//
type GroupsPermissions struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Permissions []string `json:"permissions" url:"permissions,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *GroupsPermissions) GetMethod() string {
	return "GET"
}

func (t *GroupsPermissions) GetURLPath() string {
	path := "groups/{group_id}/permissions"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *GroupsPermissions) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *GroupsPermissions) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GroupsPermissions) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GroupsPermissions) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GroupsPermissions) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
