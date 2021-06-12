package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type CommMessage struct {
	ID            int64     `json:"id"`             // The ID of the CommMessage..Example: 42
	CreatedAt     time.Time `json:"created_at"`     // The date and time this message was created.Example: 2013-03-19T21:00:00Z
	SentAt        time.Time `json:"sent_at"`        // The date and time this message was sent.Example: 2013-03-20T22:42:00Z
	WorkflowState string    `json:"workflow_state"` // The workflow state of the message. One of 'created', 'staged', 'sending', 'sent', 'bounced', 'dashboard', 'cancelled', or 'closed'.Example: sent
	From          string    `json:"from"`           // The address that was put in the 'from' field of the message.Example: notifications@example.com
	FromName      string    `json:"from_name"`      // The display name for the from address.Example: Instructure Canvas
	To            string    `json:"to"`             // The address the message was sent to:.Example: someone@example.com
	ReplyTo       string    `json:"reply_to"`       // The reply_to header of the message.Example: notifications+specialdata@example.com
	Subject       string    `json:"subject"`        // The message subject.Example: example subject line
	Body          string    `json:"body"`           // The plain text body of the message.Example: This is the body of the message
	HtmlBody      string    `json:"html_body"`      // The HTML body of the message..Example: <html><body>This is the body of the message</body></html>
}

func (t *CommMessage) HasError() error {
	var s []string
	s = []string{"created", "staged", "sending", "sent", "bounced", "dashboard", "cancelled", "closed"}
	if !string_utils.Include(s, t.WorkflowState) {
		return fmt.Errorf("expected 'workflow_state' to be one of %v", s)
	}
	return nil
}
