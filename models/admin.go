package models

type Admin struct {
	ID            int64  `json:"id"`             // The unique identifier for the account role/user assignment..Example: 1023
	Role          string `json:"role"`           // The account role assigned. This can be 'AccountAdmin' or a user-defined role created by the Roles API..Example: AccountAdmin
	User          *User  `json:"user"`           // The user the role is assigned to. See the Users API for details..
	WorkflowState string `json:"workflow_state"` // The status of the account role/user assignment..Example: deleted
}

func (t *Admin) HasError() error {
	return nil
}
