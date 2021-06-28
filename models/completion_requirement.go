package models

import (
	"fmt"

	"github.com/atomicjolt/string_utils"
)

type CompletionRequirement struct {
	Type      string `json:"type" url:"type,omitempty"`           // one of 'must_view', 'must_submit', 'must_contribute', 'min_score', 'must_mark_done'.Example: min_score
	MinScore  int64  `json:"min_score" url:"min_score,omitempty"` // minimum score required to complete (only present when type == 'min_score').Example: 10
	Completed bool   `json:"completed" url:"completed,omitempty"` // whether the calling user has met this requirement (Optional; present only if the caller is a student or if the optional parameter 'student_id' is included).Example: true
}

func (t *CompletionRequirement) HasError() error {
	var s []string
	s = []string{"must_view", "must_submit", "must_contribute", "min_score", "must_mark_done"}
	if t.Type != "" && !string_utils.Include(s, t.Type) {
		return fmt.Errorf("expected 'type' to be one of %v", s)
	}
	return nil
}
