package models

type PairingCode struct {
	UserID        int64  `json:"user_id"`        // The ID of the user..Example: 2
	Code          string `json:"code"`           // The actual code to be sent to other APIs.Example: abc123
	ExpiresAt     string `json:"expires_at"`     // When the code expires.Example: 2012-05-30T17:45:25Z
	WorkflowState string `json:"workflow_state"` // The current status of the code.Example: active
}

func (t *PairingCode) HasError() error {
	return nil
}
