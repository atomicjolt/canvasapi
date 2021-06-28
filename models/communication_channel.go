package models

import (
	"fmt"

	"github.com/atomicjolt/string_utils"
)

type CommunicationChannel struct {
	ID            int64  `json:"id" url:"id,omitempty"`                         // The ID of the communication channel..Example: 16
	Address       string `json:"address" url:"address,omitempty"`               // The address, or path, of the communication channel..Example: sheldon@caltech.example.com
	Type          string `json:"type" url:"type,omitempty"`                     // The type of communcation channel being described. Possible values are: 'email', 'push', 'sms', or 'twitter'. This field determines the type of value seen in 'address'..Example: email
	Position      int64  `json:"position" url:"position,omitempty"`             // The position of this communication channel relative to the user's other channels when they are ordered..Example: 1
	UserID        int64  `json:"user_id" url:"user_id,omitempty"`               // The ID of the user that owns this communication channel..Example: 1
	WorkflowState string `json:"workflow_state" url:"workflow_state,omitempty"` // The current state of the communication channel. Possible values are: 'unconfirmed' or 'active'..Example: active
}

func (t *CommunicationChannel) HasError() error {
	var s []string
	s = []string{"email", "push", "sms", "twitter"}
	if t.Type != "" && !string_utils.Include(s, t.Type) {
		return fmt.Errorf("expected 'type' to be one of %v", s)
	}
	s = []string{"unconfirmed", "active"}
	if t.WorkflowState != "" && !string_utils.Include(s, t.WorkflowState) {
		return fmt.Errorf("expected 'workflow_state' to be one of %v", s)
	}
	return nil
}
