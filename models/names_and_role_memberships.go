package models

type NamesAndRoleMemberships struct {
	ID      string                    `json:"id" url:"id,omitempty"`           // Invocation URL.Example: https://example.instructure.com/api/lti/courses/1/names_and_roles?tlid=f91ca4d8-fa84-4a9b-b08e-47d5527416b0
	Context *NamesAndRoleContext      `json:"context" url:"context,omitempty"` // The LTI Context containing the memberships.Example: 4dde05e8ca1973bcca9bffc13e1548820eee93a3, CS-101, Computer Science 101
	Members []*NamesAndRoleMembership `json:"members" url:"members,omitempty"` // A list of NamesAndRoleMembership.Example: {'status'=>'Active', 'name'=>'Sienna Howell', 'picture'=>'https://example.instructure.com/images/messages/avatar-50.png', 'given_name'=>'Sienna', 'family_name'=>'Howell', 'email'=>'showell@school.edu', 'lis_person_sourcedid'=>'1238.8763.00', 'user_id'=>'535fa085f22b4655f48cd5a36a9215f64c062838', 'roles'=>['http://purl.imsglobal.org/vocab/lis/v2/membership#Instructor', 'http://purl.imsglobal.org/vocab/lis/v2/membership#ContentDeveloper'], 'message'=>[{'https://purl.imsglobal.org/spec/lti/claim/message_type'=>'LtiResourceLinkRequest', 'locale'=>'en', 'https://www.instructure.com/canvas_user_id'=>1, 'https://www.instructure.com/canvas_user_login_id'=>'showell@school.edu', 'https://purl.imsglobal.org/spec/lti/claim/custom'=>{'message_locale'=>'en', 'person_address_timezone'=>'America/Denver'}}]}, {'status'=>'Active', 'name'=>'Terrence Walls', 'picture'=>'https://example.instructure.com/images/messages/avatar-51.png', 'given_name'=>'Terrence', 'family_name'=>'Walls', 'email'=>'twalls@school.edu', 'lis_person_sourcedid'=>'5790.3390.11', 'user_id'=>'86157096483e6b3a50bfedc6bac902c0b20a824f', 'roles'=>['http://purl.imsglobal.org/vocab/lis/v2/membership#Learner'], 'message'=>[{'https://purl.imsglobal.org/spec/lti/claim/message_type'=>'LtiResourceLinkRequest', 'locale'=>'de', 'https://www.instructure.com/canvas_user_id'=>2, 'https://www.instructure.com/canvas_user_login_id'=>'twalls@school.edu', 'https://purl.imsglobal.org/spec/lti/claim/custom'=>{'message_locale'=>'en', 'person_address_timezone'=>'Europe/Berlin'}}]}
}

func (t *NamesAndRoleMemberships) HasErrors() error {
	return nil
}
