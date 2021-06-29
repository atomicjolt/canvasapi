package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// CreateUserLogin Create a new login for an existing user in the given account.
// https://canvas.instructure.com/doc/api/logins.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
// Form Parameters:
// # Form.User.ID (Required) The ID of the user to create the login for.
// # Form.Login.UniqueID (Required) The unique ID for the new login.
// # Form.Login.Password (Optional) The new login's password.
// # Form.Login.SISUserID (Optional) SIS ID for the login. To set this parameter, the caller must be able to
//    manage SIS permissions on the account.
// # Form.Login.IntegrationID (Optional) Integration ID for the login. To set this parameter, the caller must be able to
//    manage SIS permissions on the account. The Integration ID is a secondary
//    identifier useful for more complex SIS integrations.
// # Form.Login.AuthenticationProviderID (Optional) The authentication provider this login is associated with. Logins
//    associated with a specific provider can only be used with that provider.
//    Legacy providers (LDAP, CAS, SAML) will search for logins associated with
//    them, or unassociated logins. New providers will only search for logins
//    explicitly associated with them. This can be the integer ID of the
//    provider, or the type of the provider (in which case, it will find the
//    first matching provider).
//
type CreateUserLogin struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		User struct {
			ID string `json:"id" url:"id,omitempty"` //  (Required)
		} `json:"user" url:"user,omitempty"`

		Login struct {
			UniqueID                 string `json:"unique_id" url:"unique_id,omitempty"`                                   //  (Required)
			Password                 string `json:"password" url:"password,omitempty"`                                     //  (Optional)
			SISUserID                string `json:"sis_user_id" url:"sis_user_id,omitempty"`                               //  (Optional)
			IntegrationID            string `json:"integration_id" url:"integration_id,omitempty"`                         //  (Optional)
			AuthenticationProviderID string `json:"authentication_provider_id" url:"authentication_provider_id,omitempty"` //  (Optional)
		} `json:"login" url:"login,omitempty"`
	} `json:"form"`
}

func (t *CreateUserLogin) GetMethod() string {
	return "POST"
}

func (t *CreateUserLogin) GetURLPath() string {
	path := "accounts/{account_id}/logins"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *CreateUserLogin) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateUserLogin) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateUserLogin) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateUserLogin) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if t.Form.User.ID == "" {
		errs = append(errs, "'Form.User.ID' is required")
	}
	if t.Form.Login.UniqueID == "" {
		errs = append(errs, "'Form.Login.UniqueID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateUserLogin) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
