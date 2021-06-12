package models

type NamesAndRoleMembership struct {
	Status             string                 `json:"status"`               // Membership state.Example: Active
	Name               string                 `json:"name"`                 // Member's full name. Only included if tool privacy level is `public` or `name_only`..Example: Sienna Howell
	Picture            string                 `json:"picture"`              // URL to the member's avatar. Only included if tool privacy level is `public`..Example: https://example.instructure.com/images/messages/avatar-50.png
	GivenName          string                 `json:"given_name"`           // Member's 'first' name. Only included if tool privacy level is `public` or `name_only`..Example: Sienna
	FamilyName         string                 `json:"family_name"`          // Member's 'last' name. Only included if tool privacy level is `public` or `name_only`..Example: Howell
	Email              string                 `json:"email"`                // Member's email address. Only included if tool privacy level is `public` or `email_only`..Example: showell@school.edu
	LisPersonSourcedid string                 `json:"lis_person_sourcedid"` // Member's primary SIS identifier. Only included if tool privacy level is `public` or `name_only`..Example: 1238.8763.00
	UserID             string                 `json:"user_id"`              // Member's unique LTI identifier..Example: 535fa085f22b4655f48cd5a36a9215f64c062838
	Roles              []string               `json:"roles"`                // Member's roles in the current Context, expressed as LTI/LIS URNs..Example: http://purl.imsglobal.org/vocab/lis/v2/membership#Instructor, http://purl.imsglobal.org/vocab/lis/v2/membership#ContentDeveloper
	Message            []*NamesAndRoleMessage `json:"message"`              // Only present when the request specifies a `rlid` query parameter. Contains additional attributes which would appear in the LTI launch message were this member to click the link referenced by the `rlid` query parameter.Example: {'https://purl.imsglobal.org/spec/lti/claim/message_type'=>'LtiResourceLinkRequest', 'locale'=>'en', 'https://www.instructure.com/canvas_user_id'=>1, 'https://www.instructure.com/canvas_user_login_id'=>'showell@school.edu', 'https://purl.imsglobal.org/spec/lti/claim/custom'=>{'message_locale'=>'en', 'person_address_timezone'=>'America/Denver'}}
}

func (t *NamesAndRoleMembership) HasError() error {
	return nil
}
