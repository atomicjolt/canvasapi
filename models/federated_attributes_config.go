package models

type FederatedAttributesConfig struct {
	AdminRoles    string `json:"admin_roles" url:"admin_roles,omitempty"`       // A comma separated list of role names to grant to the user. Note that these only apply at the root account level, and not sub-accounts. If the attribute is not marked for provisioning only, the user will also be removed from any other roles they currently hold that are not still specified by the IdP..
	DisplayName   string `json:"display_name" url:"display_name,omitempty"`     // The full display name of the user.
	Email         string `json:"email" url:"email,omitempty"`                   // The user's e-mail address.
	GivenName     string `json:"given_name" url:"given_name,omitempty"`         // The first, or given, name of the user.
	IntegrationID string `json:"integration_id" url:"integration_id,omitempty"` // The secondary unique identifier for SIS purposes.
	Locale        string `json:"locale" url:"locale,omitempty"`                 // The user's preferred locale/language.
	Name          string `json:"name" url:"name,omitempty"`                     // The full name of the user.
	SISUserID     string `json:"sis_user_id" url:"sis_user_id,omitempty"`       // The unique SIS identifier.
	SortableName  string `json:"sortable_name" url:"sortable_name,omitempty"`   // The full name of the user for sorting purposes.
	Surname       string `json:"surname" url:"surname,omitempty"`               // The surname, or last name, of the user.
	Timezone      string `json:"timezone" url:"timezone,omitempty"`             // The user's preferred time zone.
}

func (t *FederatedAttributesConfig) HasErrors() error {
	return nil
}
