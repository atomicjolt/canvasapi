package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// MakeAccountAdmin Flag an existing user as an admin within the account.
// https://canvas.instructure.com/doc/api/admins.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Form Parameters:
// # UserID (Required) The id of the user to promote.
// # Role (Optional) [DEPRECATED] The user's admin relationship with the account will be
//    created with the given role. Defaults to 'AccountAdmin'.
// # RoleID (Optional) The user's admin relationship with the account will be created with the
//    given role. Defaults to the built-in role for 'AccountAdmin'.
// # SendConfirmation (Optional) Send a notification email to
//    the new admin if true. Default is true.
//
type MakeAccountAdmin struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`

	Form struct {
		UserID           int64  `json:"user_id"`           //  (Required)
		Role             string `json:"role"`              //  (Optional)
		RoleID           int64  `json:"role_id"`           //  (Optional)
		SendConfirmation bool   `json:"send_confirmation"` //  (Optional)
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

func (t *MakeAccountAdmin) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *MakeAccountAdmin) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
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
