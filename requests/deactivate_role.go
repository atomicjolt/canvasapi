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

// DeactivateRole Deactivates a custom role.  This hides it in the user interface and prevents it
// from being assigned to new users.  Existing users assigned to the role will
// continue to function with the same permissions they had previously.
// Built-in roles cannot be deactivated.
// https://canvas.instructure.com/doc/api/roles.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
// # Path.ID (Required) ID
//
// Query Parameters:
// # Query.RoleID (Required) The unique identifier for the role
// # Query.Role (Optional) The name for the role
//
type DeactivateRole struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`

	Query struct {
		RoleID int64  `json:"role_id" url:"role_id,omitempty"` //  (Required)
		Role   string `json:"role" url:"role,omitempty"`       //  (Optional)
	} `json:"query"`
}

func (t *DeactivateRole) GetMethod() string {
	return "DELETE"
}

func (t *DeactivateRole) GetURLPath() string {
	path := "accounts/{account_id}/roles/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeactivateRole) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *DeactivateRole) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeactivateRole) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeactivateRole) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeactivateRole) Do(c *canvasapi.Canvas) (*models.Role, error) {
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
