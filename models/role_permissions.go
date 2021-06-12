package models

type RolePermissions struct {
	Enabled              bool `json:"enabled"`                // Whether the role has the permission.Example: true
	Locked               bool `json:"locked"`                 // Whether the permission is locked by this role.
	AppliesToSelf        bool `json:"applies_to_self"`        // Whether the permission applies to the account this role is in. Only present if enabled is true.Example: true
	AppliesToDescendants bool `json:"applies_to_descendants"` // Whether the permission cascades down to sub accounts of the account this role is in. Only present if enabled is true.
	Readonly             bool `json:"readonly"`               // Whether the permission can be modified in this role (i.e. whether the permission is locked by an upstream role)..
	Explicit             bool `json:"explicit"`               // Whether the value of enabled is specified explicitly by this role, or inherited from an upstream role..Example: true
	PriorDefault         bool `json:"prior_default"`          // The value that would have been inherited from upstream if the role had not explicitly set a value. Only present if explicit is true..
}

func (t *RolePermissions) HasError() error {
	return nil
}
