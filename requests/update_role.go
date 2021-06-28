package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// UpdateRole Update permissions for an existing role.
//
// Recognized roles are:
// * TeacherEnrollment
// * StudentEnrollment
// * TaEnrollment
// * ObserverEnrollment
// * DesignerEnrollment
// * AccountAdmin
// * Any previously created custom role
// https://canvas.instructure.com/doc/api/roles.html
//
// Path Parameters:
// # AccountID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # Label (Optional) The label for the role. Can only change the label of a custom role that belongs directly to the account.
// # Permissions (Optional) no description
// # Permissions (Optional) These arguments are described in the documentation for the
//    {api:RoleOverridesController#add_role add_role method}.
// # Permissions (Optional) If the value is 1, permission <X> applies to the account this role is in.
//    The default value is 1. Must be true if applies_to_descendants is false.
//    This value is only returned if enabled is true.
// # Permissions (Optional) If the value is 1, permission <X> cascades down to sub accounts of the
//    account this role is in. The default value is 1.  Must be true if
//    applies_to_self is false.This value is only returned if enabled is true.
//
type UpdateRole struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`

	Form struct {
		Label       string `json:"label" url:"label,omitempty"` //  (Optional)
		Permissions map[string]UpdateRolePermissions
	} `json:"form"`
}

func (t *UpdateRole) GetMethod() string {
	return "PUT"
}

func (t *UpdateRole) GetURLPath() string {
	path := "accounts/{account_id}/roles/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateRole) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateRole) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateRole) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateRole) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateRole) Do(c *canvasapi.Canvas) (*models.Role, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Role{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

type UpdateRolePermissions struct {
	Explicit             bool `json:"explicit" url:"explicit,omitempty"`                             //  (Optional)
	Enabled              bool `json:"enabled" url:"enabled,omitempty"`                               //  (Optional)
	AppliesToSelf        bool `json:"applies_to_self" url:"applies_to_self,omitempty"`               //  (Optional)
	AppliesToDescendants bool `json:"applies_to_descendants" url:"applies_to_descendants,omitempty"` //  (Optional)
}
