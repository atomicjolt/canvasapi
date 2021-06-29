package models

type SSOSettings struct {
	LoginHandleName   string `json:"login_handle_name" url:"login_handle_name,omitempty"`     // The label used for unique login identifiers..Example: Username
	ChangePasswordUrl string `json:"change_password_url" url:"change_password_url,omitempty"` // The url to redirect users to for password resets. Leave blank for default Canvas behavior.Example: https://example.com/reset_password
	AuthDiscoveryUrl  string `json:"auth_discovery_url" url:"auth_discovery_url,omitempty"`   // If a discovery url is set, canvas will forward all users to that URL when they need to be authenticated. That page will need to then help the user figure out where they need to go to log in. If no discovery url is configured, the first configuration will be used to attempt to authenticate the user..Example: https://example.com/which_account
	UnknownUserUrl    string `json:"unknown_user_url" url:"unknown_user_url,omitempty"`       // If an unknown user url is set, Canvas will forward to that url when a service authenticates a user, but that user does not exist in Canvas. The default behavior is to present an error..Example: https://example.com/register_for_canvas
}

func (t *SSOSettings) HasErrors() error {
	return nil
}
