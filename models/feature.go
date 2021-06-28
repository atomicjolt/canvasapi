package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type Feature struct {
	Feature     string `json:"feature" url:"feature,omitempty"`           // The symbolic name of the feature, used in FeatureFlags.Example: fancy_wickets
	DisplayName string `json:"display_name" url:"display_name,omitempty"` // The user-visible name of the feature.Example: Fancy Wickets
	AppliesTo   string `json:"applies_to" url:"applies_to,omitempty"`     // The type of object the feature applies to (RootAccount, Account, Course, or User):
	// * RootAccount features may only be controlled by flags on root accounts.
	// * Account features may be controlled by flags on accounts and their parent accounts.
	// * Course features may be controlled by flags on courses and their parent accounts.
	// * User features may be controlled by flags on users and site admin only..Example: Course
	EnableAt        time.Time    `json:"enable_at" url:"enable_at,omitempty"`                 // The date this feature will be globally enabled, or null if this is not planned. (This information is subject to change.).Example: 2014-01-01T00:00:00Z
	FeatureFlag     *FeatureFlag `json:"feature_flag" url:"feature_flag,omitempty"`           // The FeatureFlag that applies to the caller.Example: fancy_wickets, allowed
	RootOptIn       bool         `json:"root_opt_in" url:"root_opt_in,omitempty"`             // If true, a feature that is 'allowed' globally will be 'off' by default in root accounts. Otherwise, root accounts inherit the global 'allowed' setting, which allows sub-accounts and courses to turn features on with no root account action..Example: true
	Beta            bool         `json:"beta" url:"beta,omitempty"`                           // Whether the feature is a beta feature. If true, the feature may not be fully polished and may be subject to change in the future..Example: true
	Autoexpand      bool         `json:"autoexpand" url:"autoexpand,omitempty"`               // Whether the details of the feature are autoexpanded on page load vs. the user clicking to expand..Example: true
	Development     bool         `json:"development" url:"development,omitempty"`             // Whether the feature is in active development. Features in this state are only visible in test and beta instances and are not yet available for production use..
	ReleaseNotesUrl string       `json:"release_notes_url" url:"release_notes_url,omitempty"` // A URL to the release notes describing the feature.Example: http://canvas.example.com/release_notes#fancy_wickets
}

func (t *Feature) HasError() error {
	var s []string
	s = []string{"Course", "RootAccount", "Account", "User"}
	if t.AppliesTo != "" && !string_utils.Include(s, t.AppliesTo) {
		return fmt.Errorf("expected 'applies_to' to be one of %v", s)
	}
	return nil
}
