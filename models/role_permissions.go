package models

type RolePermissions struct {
	Enabled              bool `json:"enabled" url:"enabled,omitempty"`                               // Whether the role has the permission.Example: true
	Locked               bool `json:"locked" url:"locked,omitempty"`                                 // Whether the permission is locked by this role.
	AppliesToSelf        bool `json:"applies_to_self" url:"applies_to_self,omitempty"`               // Whether the permission applies to the account this role is in. Only present if enabled is true.Example: true
	AppliesToDescendants bool `json:"applies_to_descendants" url:"applies_to_descendants,omitempty"` // Whether the permission cascades down to sub accounts of the account this role is in. Only present if enabled is true.
	Readonly             bool `json:"readonly" url:"readonly,omitempty"`                             // Whether the permission can be modified in this role (i.e. whether the permission is locked by an upstream role)..
	Explicit             bool `json:"explicit" url:"explicit,omitempty"`                             // Whether the value of enabled is specified explicitly by this role, or inherited from an upstream role..Example: true
	PriorDefault         bool `json:"prior_default" url:"prior_default,omitempty"`                   // The value that would have been inherited from upstream if the role had not explicitly set a value. Only present if explicit is true..
}

func (t *RolePermissions) HasErrors() error {
	return nil
}
