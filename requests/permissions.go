package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// Permissions Returns permission information for the calling user and the given account.
// You may use `self` as the account id to check permissions against the domain root account.
// The caller must have an account role or admin (teacher/TA/designer) enrollment in a course
// in the account.
//
// See also the {api:CoursesController#permissions Course} and {api:GroupsController#permissions Group}
// counterparts.
// https://canvas.instructure.com/doc/api/accounts.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Query Parameters:
// # Permissions (Optional) List of permissions to check against the authenticated user.
//    Permission names are documented in the {api:RoleOverridesController#add_role Create a role} endpoint.
//
type Permissions struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Permissions []string `json:"permissions" url:"permissions,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *Permissions) GetMethod() string {
	return "GET"
}

func (t *Permissions) GetURLPath() string {
	path := "accounts/{account_id}/permissions"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *Permissions) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *Permissions) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *Permissions) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *Permissions) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *Permissions) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
