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

// CreateUser Create and return a new user and pseudonym for an account.
//
// [DEPRECATED (for self-registration only)] If you don't have the "Modify
// login details for users" permission, but self-registration is enabled
// on the account, you can still use this endpoint to register new users.
// Certain fields will be required, and others will be ignored (see below).
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Form Parameters:
// # User (Optional) The full name of the user. This name will be used by teacher for grading.
//    Required if this is a self-registration.
// # User (Optional) User's name as it will be displayed in discussions, messages, and comments.
// # User (Optional) User's name as used to sort alphabetically in lists.
// # User (Optional) The time zone for the user. Allowed time zones are
//    {http://www.iana.org/time-zones IANA time zones} or friendlier
//    {http://api.rubyonrails.org/classes/ActiveSupport/TimeZone.html Ruby on Rails time zones}.
// # User (Optional) The user's preferred language, from the list of languages Canvas supports.
//    This is in RFC-5646 format.
// # User (Optional) Whether the user accepts the terms of use. Required if this is a
//    self-registration and this canvas instance requires users to accept
//    the terms (on by default).
//
//    If this is true, it will mark the user as having accepted the terms of use.
// # User (Optional) Automatically mark the user as registered.
//
//    If this is true, it is recommended to set <tt>"pseudonym[send_confirmation]"</tt> to true as well.
//    Otherwise, the user will not receive any messages about their account creation.
//
//    The users communication channel confirmation can be skipped by setting
//    <tt>"communication_channel[skip_confirmation]"</tt> to true as well.
// # Pseudonym (Required) User's login ID. If this is a self-registration, it must be a valid
//    email address.
// # Pseudonym (Optional) User's password. Cannot be set during self-registration.
// # Pseudonym (Optional) SIS ID for the user's account. To set this parameter, the caller must be
//    able to manage SIS permissions.
// # Pseudonym (Optional) Integration ID for the login. To set this parameter, the caller must be able to
//    manage SIS permissions. The Integration ID is a secondary
//    identifier useful for more complex SIS integrations.
// # Pseudonym (Optional) Send user notification of account creation if true.
//    Automatically set to true during self-registration.
// # Pseudonym (Optional) Send user a self-registration style email if true.
//    Setting it means the users will get a notification asking them
//    to "complete the registration process" by clicking it, setting
//    a password, and letting them in.  Will only be executed on
//    if the user does not need admin approval.
//    Defaults to false unless explicitly provided.
// # Pseudonym (Optional) The authentication provider this login is associated with. Logins
//    associated with a specific provider can only be used with that provider.
//    Legacy providers (LDAP, CAS, SAML) will search for logins associated with
//    them, or unassociated logins. New providers will only search for logins
//    explicitly associated with them. This can be the integer ID of the
//    provider, or the type of the provider (in which case, it will find the
//    first matching provider).
// # CommunicationChannel (Optional) The communication channel type, e.g. 'email' or 'sms'.
// # CommunicationChannel (Optional) The communication channel address, e.g. the user's email address.
// # CommunicationChannel (Optional) Only valid for account admins. If true, returns the new user account
//    confirmation URL in the response.
// # CommunicationChannel (Optional) Only valid for site admins and account admins making requests; If true, the channel is
//    automatically validated and no confirmation email or SMS is sent.
//    Otherwise, the user must respond to a confirmation message to confirm the
//    channel.
//
//    If this is true, it is recommended to set <tt>"pseudonym[send_confirmation]"</tt> to true as well.
//    Otherwise, the user will not receive any messages about their account creation.
// # ForceValidations (Optional) If true, validations are performed on the newly created user (and their associated pseudonym)
//    even if the request is made by a privileged user like an admin. When set to false,
//    or not included in the request parameters, any newly created users are subject to
//    validations unless the request is made by a user with a 'manage_user_logins' right.
//    In which case, certain validations such as 'require_acceptance_of_terms' and
//    'require_presence_of_name' are not enforced. Use this parameter to return helpful json
//    errors while building users with an admin request.
// # EnableSISReactivation (Optional) When true, will first try to re-activate a deleted user with matching sis_user_id if possible.
//    This is commonly done with user[skip_registration] and communication_channel[skip_confirmation]
//    so that the default communication_channel is also restored.
// # Destination (Optional) If you're setting the password for the newly created user, you can provide this param
//    with a valid URL pointing into this Canvas installation, and the response will include
//    a destination field that's a URL that you can redirect a browser to and have the newly
//    created user automatically logged in. The URL is only valid for a short time, and must
//    match the domain this request is directed to, and be for a well-formed path that Canvas
//    can recognize.
// # InitialEnrollmentType (Optional) `observer` if doing a self-registration with a pairing code. This allows setting the
//    password during user creation.
// # PairingCode (Optional) If provided and valid, will link the new user as an observer to the student's whose
//    pairing code is given.
//
type CreateUser struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`

	Form struct {
		User struct {
			Name             string `json:"name"`              //  (Optional)
			ShortName        string `json:"short_name"`        //  (Optional)
			SortableName     string `json:"sortable_name"`     //  (Optional)
			TimeZone         string `json:"time_zone"`         //  (Optional)
			Locale           string `json:"locale"`            //  (Optional)
			TermsOfUse       bool   `json:"terms_of_use"`      //  (Optional)
			SkipRegistration bool   `json:"skip_registration"` //  (Optional)
		} `json:"user"`

		Pseudonym struct {
			UniqueID                 string `json:"unique_id"`                  //  (Required)
			Password                 string `json:"password"`                   //  (Optional)
			SISUserID                string `json:"sis_user_id"`                //  (Optional)
			IntegrationID            string `json:"integration_id"`             //  (Optional)
			SendConfirmation         bool   `json:"send_confirmation"`          //  (Optional)
			ForceSelfRegistration    bool   `json:"force_self_registration"`    //  (Optional)
			AuthenticationProviderID string `json:"authentication_provider_id"` //  (Optional)
		} `json:"pseudonym"`

		CommunicationChannel struct {
			Type             string `json:"type"`              //  (Optional)
			Address          string `json:"address"`           //  (Optional)
			ConfirmationUrl  bool   `json:"confirmation_url"`  //  (Optional)
			SkipConfirmation bool   `json:"skip_confirmation"` //  (Optional)
		} `json:"communication_channel"`

		ForceValidations      bool   `json:"force_validations"`       //  (Optional)
		EnableSISReactivation bool   `json:"enable_sis_reactivation"` //  (Optional)
		Destination           string `json:"destination"`             //  (Optional)
		InitialEnrollmentType string `json:"initial_enrollment_type"` //  (Optional)
		PairingCode           struct {
			Code string `json:"code"` //  (Optional)
		} `json:"pairing_code"`
	} `json:"form"`
}

func (t *CreateUser) GetMethod() string {
	return "POST"
}

func (t *CreateUser) GetURLPath() string {
	path := "accounts/{account_id}/users"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *CreateUser) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateUser) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateUser) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Form.Pseudonym.UniqueID == "" {
		errs = append(errs, "'Pseudonym' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateUser) Do(c *canvasapi.Canvas) (*models.User, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.User{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
