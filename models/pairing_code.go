package models

type PairingCode struct {
	UserID        int64  `json:"user_id" url:"user_id,omitempty"`               // The ID of the user..Example: 2
	Code          string `json:"code" url:"code,omitempty"`                     // The actual code to be sent to other APIs.Example: abc123
	ExpiresAt     string `json:"expires_at" url:"expires_at,omitempty"`         // When the code expires.Example: 2012-05-30T17:45:25Z
	WorkflowState string `json:"workflow_state" url:"workflow_state,omitempty"` // The current status of the code.Example: active
}

func (t *PairingCode) HasErrors() error {
	return nil
}
