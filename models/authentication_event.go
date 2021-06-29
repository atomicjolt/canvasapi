package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type AuthenticationEvent struct {
	CreatedAt   time.Time `json:"created_at" url:"created_at,omitempty"`     // timestamp of the event.Example: 2012-07-19T15:00:00-06:00
	EventType   string    `json:"event_type" url:"event_type,omitempty"`     // authentication event type ('login' or 'logout').Example: login
	PseudonymID int64     `json:"pseudonym_id" url:"pseudonym_id,omitempty"` // ID of the pseudonym (login) associated with the event.Example: 9478
	AccountID   int64     `json:"account_id" url:"account_id,omitempty"`     // ID of the account associated with the event. will match the account_id in the associated pseudonym..Example: 2319
	UserID      int64     `json:"user_id" url:"user_id,omitempty"`           // ID of the user associated with the event will match the user_id in the associated pseudonym..Example: 362
}

func (t *AuthenticationEvent) HasErrors() error {
	var s []string
	errs := []string{}
	s = []string{"login", "logout"}
	if t.EventType != "" && !string_utils.Include(s, t.EventType) {
		errs = append(errs, fmt.Sprintf("expected 'EventType' to be one of %v", s))
	}
	return nil
}
