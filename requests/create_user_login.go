package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// CreateUserLogin Create a new login for an existing user in the given account.
// https://canvas.instructure.com/doc/api/logins.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Form Parameters:
// # User (Required) The ID of the user to create the login for.
// # Login (Required) The unique ID for the new login.
// # Login (Optional) The new login's password.
// # Login (Optional) SIS ID for the login. To set this parameter, the caller must be able to
//    manage SIS permissions on the account.
// # Login (Optional) Integration ID for the login. To set this parameter, the caller must be able to
//    manage SIS permissions on the account. The Integration ID is a secondary
//    identifier useful for more complex SIS integrations.
// # Login (Optional) The authentication provider this login is associated with. Logins
//    associated with a specific provider can only be used with that provider.
//    Legacy providers (LDAP, CAS, SAML) will search for logins associated with
//    them, or unassociated logins. New providers will only search for logins
//    explicitly associated with them. This can be the integer ID of the
//    provider, or the type of the provider (in which case, it will find the
//    first matching provider).
//
type CreateUserLogin struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`

	Form struct {
		User struct {
			ID string `json:"id"` //  (Required)
		} `json:"user"`

		Login struct {
			UniqueID                 string `json:"unique_id"`                  //  (Required)
			Password                 string `json:"password"`                   //  (Optional)
			SISUserID                string `json:"sis_user_id"`                //  (Optional)
			IntegrationID            string `json:"integration_id"`             //  (Optional)
			AuthenticationProviderID string `json:"authentication_provider_id"` //  (Optional)
		} `json:"login"`
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

func (t *CreateUserLogin) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateUserLogin) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Form.User.ID == "" {
		errs = append(errs, "'User' is required")
	}
	if t.Form.Login.UniqueID == "" {
		errs = append(errs, "'Login' is required")
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
