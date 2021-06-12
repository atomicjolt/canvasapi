package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type AuthenticationEvent struct {
	CreatedAt   time.Time `json:"created_at"`   // timestamp of the event.Example: 2012-07-19T15:00:00-06:00
	EventType   string    `json:"event_type"`   // authentication event type ('login' or 'logout').Example: login
	PseudonymID int64     `json:"pseudonym_id"` // ID of the pseudonym (login) associated with the event.Example: 9478
	AccountID   int64     `json:"account_id"`   // ID of the account associated with the event. will match the account_id in the associated pseudonym..Example: 2319
	UserID      int64     `json:"user_id"`      // ID of the user associated with the event will match the user_id in the associated pseudonym..Example: 362
}

func (t *AuthenticationEvent) HasError() error {
	var s []string
	s = []string{"login", "logout"}
	if !string_utils.Include(s, t.EventType) {
		return fmt.Errorf("expected 'event_type' to be one of %v", s)
	}
	return nil
}
