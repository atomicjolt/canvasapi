package models

type Role struct {
	Label         string   `json:"label" url:"label,omitempty"`                   // The label of the role..Example: New Role
	Role          string   `json:"role" url:"role,omitempty"`                     // The label of the role. (Deprecated alias for 'label').Example: New Role
	BaseRoleType  string   `json:"base_role_type" url:"base_role_type,omitempty"` // The role type that is being used as a base for this role. For account-level roles, this is 'AccountMembership'. For course-level roles, it is an enrollment type..Example: AccountMembership
	Account       *Account `json:"account" url:"account,omitempty"`               // JSON representation of the account the role is in..Example: 1019, CGNU, 73, 1, cgnu
	WorkflowState string   `json:"workflow_state" url:"workflow_state,omitempty"` // The state of the role: 'active', 'inactive', or 'built_in'.Example: active
	Permissions   string   `json:"permissions" url:"permissions,omitempty"`       // A dictionary of permissions keyed by name (see permissions input parameter in the 'Create a role' API)..Example: {'enabled'=>true, 'locked'=>false, 'readonly'=>false, 'explicit'=>true, 'prior_default'=>false}, {'enabled'=>true, 'locked'=>true, 'readonly'=>true, 'explicit'=>false}, {'enabled'=>false, 'locked'=>true, 'readonly'=>false, 'explicit'=>true, 'prior_default'=>false}, {'enabled'=>true, 'locked'=>false, 'readonly'=>false, 'explicit'=>false}
}

func (t *Role) HasError() error {
	return nil
}
