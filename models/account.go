package models

type Account struct {
	ID              int64  `json:"id" url:"id,omitempty"`                               // the ID of the Account object.Example: 2
	Name            string `json:"name" url:"name,omitempty"`                           // The display name of the account.Example: Canvas Account
	Uuid            string `json:"uuid" url:"uuid,omitempty"`                           // The UUID of the account.Example: WvAHhY5FINzq5IyRIJybGeiXyFkG3SqHUPb7jZY5
	ParentAccountID int64  `json:"parent_account_id" url:"parent_account_id,omitempty"` // The account's parent ID, or null if this is the root account.Example: 1
	RootAccountID   int64  `json:"root_account_id" url:"root_account_id,omitempty"`     // The ID of the root account, or null if this is the root account.Example: 1
	WorkflowState   string `json:"workflow_state" url:"workflow_state,omitempty"`       // The state of the account. Can be 'active' or 'deleted'..Example: active
}

func (t *Account) HasError() error {
	return nil
}
