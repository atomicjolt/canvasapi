package models

import (
	"fmt"

	"github.com/atomicjolt/string_utils"
)

type FeatureFlag struct {
	ContextType string `json:"context_type"` // The type of object to which this flag applies (Account, Course, or User). (This field is not present if this FeatureFlag represents the global Canvas default).Example: Account
	ContextID   int64  `json:"context_id"`   // The id of the object to which this flag applies (This field is not present if this FeatureFlag represents the global Canvas default).Example: 1038
	Feature     string `json:"feature"`      // The feature this flag controls.Example: fancy_wickets
	State       string `json:"state"`        // The policy for the feature at this context.  can be 'off', 'allowed', 'allowed_on', or 'on'..Example: allowed
	Locked      bool   `json:"locked"`       // If set, this feature flag cannot be changed in the caller's context because the flag is set 'off' or 'on' in a higher context.
}

func (t *FeatureFlag) HasError() error {
	var s []string
	s = []string{"Course", "Account", "User"}
	if !string_utils.Include(s, t.ContextType) {
		return fmt.Errorf("expected 'context_type' to be one of %v", s)
	}
	s = []string{"off", "allowed", "allowed_on", "on"}
	if !string_utils.Include(s, t.State) {
		return fmt.Errorf("expected 'state' to be one of %v", s)
	}
	return nil
}
