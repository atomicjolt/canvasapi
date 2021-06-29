package models

import (
	"fmt"

	"github.com/atomicjolt/string_utils"
)

type Favorite struct {
	ContextID   int64  `json:"context_id" url:"context_id,omitempty"`     // The ID of the object the Favorite refers to.Example: 1170
	ContextType string `json:"context_type" url:"context_type,omitempty"` // The type of the object the Favorite refers to (currently, only 'Course' is supported).Example: Course
}

func (t *Favorite) HasErrors() error {
	var s []string
	errs := []string{}
	s = []string{"Course"}
	if t.ContextType != "" && !string_utils.Include(s, t.ContextType) {
		errs = append(errs, fmt.Sprintf("expected 'ContextType' to be one of %v", s))
	}
	return nil
}
