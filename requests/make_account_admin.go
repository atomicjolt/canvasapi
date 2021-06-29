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

// MakeAccountAdmin Flag an existing user as an admin within the account.
// https://canvas.instructure.com/doc/api/admins.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
// Form Parameters:
// # Form.UserID (Required) The id of the user to promote.
// # Form.Role (Optional) [DEPRECATED] The user's admin relationship with the account will be
//    created with the given role. Defaults to 'AccountAdmin'.
// # Form.RoleID (Optional) The user's admin relationship with the account will be created with the
//    given role. Defaults to the built-in role for 'AccountAdmin'.
// # Form.SendConfirmation (Optional) Send a notification email to
//    the new admin if true. Default is true.
//
type MakeAccountAdmin struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		UserID           int64  `json:"user_id" url:"user_id,omitempty"`                     //  (Required)
		Role             string `json:"role" url:"role,omitempty"`                           //  (Optional)
		RoleID           int64  `json:"role_id" url:"role_id,omitempty"`                     //  (Optional)
		SendConfirmation bool   `json:"send_confirmation" url:"send_confirmation,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *MakeAccountAdmin) GetMethod() string {
	return "POST"
}

func (t *MakeAccountAdmin) GetURLPath() string {
	path := "accounts/{account_id}/admins"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *MakeAccountAdmin) GetQuery() (string, error) {
	return "", nil
}

func (t *MakeAccountAdmin) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *MakeAccountAdmin) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *MakeAccountAdmin) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *MakeAccountAdmin) Do(c *canvasapi.Canvas) (*models.Admin, error) {
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
