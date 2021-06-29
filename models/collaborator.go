package models

import (
	"fmt"

	"github.com/atomicjolt/string_utils"
)

type Collaborator struct {
	ID   int64  `json:"id" url:"id,omitempty"`     // The unique user or group identifier for the collaborator..Example: 12345
	Type string `json:"type" url:"type,omitempty"` // The type of collaborator (e.g. 'user' or 'group')..Example: user
	Name string `json:"name" url:"name,omitempty"` // The name of the collaborator..Example: Don Draper
}

func (t *Collaborator) HasErrors() error {
	var s []string
	errs := []string{}
	s = []string{"user", "group"}
	if t.Type != "" && !string_utils.Include(s, t.Type) {
		errs = append(errs, fmt.Sprintf("expected 'Type' to be one of %v", s))
	}
	return nil
}
