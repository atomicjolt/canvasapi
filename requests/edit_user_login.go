package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// EditUserLogin Update an existing login for a user in the given account.
// https://canvas.instructure.com/doc/api/logins.html
//
// Path Parameters:
// # AccountID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # Login (Optional) The new unique ID for the login.
// # Login (Optional) The new password for the login. Can only be set by an admin user if admins
//    are allowed to change passwords for the account.
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
type EditUserLogin struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`

	Form struct {
		Login struct {
			UniqueID                 string `json:"unique_id" url:"unique_id,omitempty"`                                   //  (Optional)
			Password                 string `json:"password" url:"password,omitempty"`                                     //  (Optional)
			SISUserID                string `json:"sis_user_id" url:"sis_user_id,omitempty"`                               //  (Optional)
			IntegrationID            string `json:"integration_id" url:"integration_id,omitempty"`                         //  (Optional)
			AuthenticationProviderID string `json:"authentication_provider_id" url:"authentication_provider_id,omitempty"` //  (Optional)
		} `json:"login" url:"login,omitempty"`
	} `json:"form"`
}

func (t *EditUserLogin) GetMethod() string {
	return "PUT"
}

func (t *EditUserLogin) GetURLPath() string {
	path := "accounts/{account_id}/logins/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *EditUserLogin) GetQuery() (string, error) {
	return "", nil
}

func (t *EditUserLogin) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *EditUserLogin) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *EditUserLogin) HasErrors() error {
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

func (t *EditUserLogin) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
