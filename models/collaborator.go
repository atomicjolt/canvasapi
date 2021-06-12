package models

import (
	"fmt"

	"github.com/atomicjolt/string_utils"
)

type Collaborator struct {
	ID   int64  `json:"id"`   // The unique user or group identifier for the collaborator..Example: 12345
	Type string `json:"type"` // The type of collaborator (e.g. 'user' or 'group')..Example: user
	Name string `json:"name"` // The name of the collaborator..Example: Don Draper
}

func (t *Collaborator) HasError() error {
	var s []string
	s = []string{"user", "group"}
	if !string_utils.Include(s, t.Type) {
		return fmt.Errorf("expected 'type' to be one of %v", s)
	}
	return nil
}
