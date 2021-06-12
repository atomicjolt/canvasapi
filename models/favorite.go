package models

import (
	"fmt"

	"github.com/atomicjolt/string_utils"
)

type Favorite struct {
	ContextID   int64  `json:"context_id"`   // The ID of the object the Favorite refers to.Example: 1170
	ContextType string `json:"context_type"` // The type of the object the Favorite refers to (currently, only 'Course' is supported).Example: Course
}

func (t *Favorite) HasError() error {
	var s []string
	s = []string{"Course"}
	if !string_utils.Include(s, t.ContextType) {
		return fmt.Errorf("expected 'context_type' to be one of %v", s)
	}
	return nil
}
