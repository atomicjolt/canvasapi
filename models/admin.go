package models

type Admin struct {
	ID            int64  `json:"id" url:"id,omitempty"`                         // The unique identifier for the account role/user assignment..Example: 1023
	Role          string `json:"role" url:"role,omitempty"`                     // The account role assigned. This can be 'AccountAdmin' or a user-defined role created by the Roles API..Example: AccountAdmin
	User          *User  `json:"user" url:"user,omitempty"`                     // The user the role is assigned to. See the Users API for details..
	WorkflowState string `json:"workflow_state" url:"workflow_state,omitempty"` // The status of the account role/user assignment..Example: deleted
}

func (t *Admin) HasError() error {
	return nil
}
