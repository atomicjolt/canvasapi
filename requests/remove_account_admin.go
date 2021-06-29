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

// RemoveAccountAdmin Remove the rights associated with an account admin role from a user.
// https://canvas.instructure.com/doc/api/admins.html
//
// Path Parameters:
// # AccountID (Required) ID
// # UserID (Required) ID
//
// Query Parameters:
// # Role (Optional) [DEPRECATED] Account role to remove from the user. Defaults to
//    'AccountAdmin'. Any other account role must be specified explicitly.
// # RoleID (Optional) The user's admin relationship with the account will be created with the
//    given role. Defaults to the built-in role for 'AccountAdmin'.
//
type RemoveAccountAdmin struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		UserID    string `json:"user_id" url:"user_id,omitempty"`       //  (Required)
	} `json:"path"`

	Query struct {
		Role   string `json:"role" url:"role,omitempty"`       //  (Optional)
		RoleID int64  `json:"role_id" url:"role_id,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *RemoveAccountAdmin) GetMethod() string {
	return "DELETE"
}

func (t *RemoveAccountAdmin) GetURLPath() string {
	path := "accounts/{account_id}/admins/{user_id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *RemoveAccountAdmin) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *RemoveAccountAdmin) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *RemoveAccountAdmin) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *RemoveAccountAdmin) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *RemoveAccountAdmin) Do(c *canvasapi.Canvas) (*models.Admin, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Admin{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
