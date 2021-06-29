package models

type NamesAndRoleMembership struct {
	Status             string                 `json:"status" url:"status,omitempty"`                             // Membership state.Example: Active
	Name               string                 `json:"name" url:"name,omitempty"`                                 // Member's full name. Only included if tool privacy level is `public` or `name_only`..Example: Sienna Howell
	Picture            string                 `json:"picture" url:"picture,omitempty"`                           // URL to the member's avatar. Only included if tool privacy level is `public`..Example: https://example.instructure.com/images/messages/avatar-50.png
	GivenName          string                 `json:"given_name" url:"given_name,omitempty"`                     // Member's 'first' name. Only included if tool privacy level is `public` or `name_only`..Example: Sienna
	FamilyName         string                 `json:"family_name" url:"family_name,omitempty"`                   // Member's 'last' name. Only included if tool privacy level is `public` or `name_only`..Example: Howell
	Email              string                 `json:"email" url:"email,omitempty"`                               // Member's email address. Only included if tool privacy level is `public` or `email_only`..Example: showell@school.edu
	LisPersonSourcedid string                 `json:"lis_person_sourcedid" url:"lis_person_sourcedid,omitempty"` // Member's primary SIS identifier. Only included if tool privacy level is `public` or `name_only`..Example: 1238.8763.00
	UserID             string                 `json:"user_id" url:"user_id,omitempty"`                           // Member's unique LTI identifier..Example: 535fa085f22b4655f48cd5a36a9215f64c062838
	Roles              []string               `json:"roles" url:"roles,omitempty"`                               // Member's roles in the current Context, expressed as LTI/LIS URNs..Example: http://purl.imsglobal.org/vocab/lis/v2/membership#Instructor, http://purl.imsglobal.org/vocab/lis/v2/membership#ContentDeveloper
	Message            []*NamesAndRoleMessage `json:"message" url:"message,omitempty"`                           // Only present when the request specifies a `rlid` query parameter. Contains additional attributes which would appear in the LTI launch message were this member to click the link referenced by the `rlid` query parameter.Example: {'https://purl.imsglobal.org/spec/lti/claim/message_type'=>'LtiResourceLinkRequest', 'locale'=>'en', 'https://www.instructure.com/canvas_user_id'=>1, 'https://www.instructure.com/canvas_user_login_id'=>'showell@school.edu', 'https://purl.imsglobal.org/spec/lti/claim/custom'=>{'message_locale'=>'en', 'person_address_timezone'=>'America/Denver'}}
}

func (t *NamesAndRoleMembership) HasErrors() error {
	return nil
}
