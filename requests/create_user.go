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
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		User struct {
			Name             string `json:"name" url:"name,omitempty"`                           //  (Optional)
			ShortName        string `json:"short_name" url:"short_name,omitempty"`               //  (Optional)
			SortableName     string `json:"sortable_name" url:"sortable_name,omitempty"`         //  (Optional)
			TimeZone         string `json:"time_zone" url:"time_zone,omitempty"`                 //  (Optional)
			Locale           string `json:"locale" url:"locale,omitempty"`                       //  (Optional)
			TermsOfUse       bool   `json:"terms_of_use" url:"terms_of_use,omitempty"`           //  (Optional)
			SkipRegistration bool   `json:"skip_registration" url:"skip_registration,omitempty"` //  (Optional)
		} `json:"user" url:"user,omitempty"`

		Pseudonym struct {
			UniqueID                 string `json:"unique_id" url:"unique_id,omitempty"`                                   //  (Required)
			Password                 string `json:"password" url:"password,omitempty"`                                     //  (Optional)
			SISUserID                string `json:"sis_user_id" url:"sis_user_id,omitempty"`                               //  (Optional)
			IntegrationID            string `json:"integration_id" url:"integration_id,omitempty"`                         //  (Optional)
			SendConfirmation         bool   `json:"send_confirmation" url:"send_confirmation,omitempty"`                   //  (Optional)
			ForceSelfRegistration    bool   `json:"force_self_registration" url:"force_self_registration,omitempty"`       //  (Optional)
			AuthenticationProviderID string `json:"authentication_provider_id" url:"authentication_provider_id,omitempty"` //  (Optional)
		} `json:"pseudonym" url:"pseudonym,omitempty"`

		CommunicationChannel struct {
			Type             string `json:"type" url:"type,omitempty"`                           //  (Optional)
			Address          string `json:"address" url:"address,omitempty"`                     //  (Optional)
			ConfirmationUrl  bool   `json:"confirmation_url" url:"confirmation_url,omitempty"`   //  (Optional)
			SkipConfirmation bool   `json:"skip_confirmation" url:"skip_confirmation,omitempty"` //  (Optional)
		} `json:"communication_channel" url:"communication_channel,omitempty"`

		ForceValidations      bool   `json:"force_validations" url:"force_validations,omitempty"`             //  (Optional)
		EnableSISReactivation bool   `json:"enable_sis_reactivation" url:"enable_sis_reactivation,omitempty"` //  (Optional)
		Destination           string `json:"destination" url:"destination,omitempty"`                         //  (Optional)
		InitialEnrollmentType string `json:"initial_enrollment_type" url:"initial_enrollment_type,omitempty"` //  (Optional)
		PairingCode           struct {
			Code string `json:"code" url:"code,omitempty"` //  (Optional)
		} `json:"pairing_code" url:"pairing_code,omitempty"`
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

func (t *CreateUser) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateUser) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
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
