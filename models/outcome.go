package models

import (
	"fmt"

	"github.com/atomicjolt/string_utils"
)

type Outcome struct {
	ID                   int64           `json:"id" url:"id,omitempty"`                                         // the ID of the outcome.Example: 1
	Url                  string          `json:"url" url:"url,omitempty"`                                       // the URL for fetching/updating the outcome. should be treated as opaque.Example: /api/v1/outcomes/1
	ContextID            int64           `json:"context_id" url:"context_id,omitempty"`                         // the context owning the outcome. may be null for global outcomes.Example: 1
	ContextType          string          `json:"context_type" url:"context_type,omitempty"`                     // Example: Account
	Title                string          `json:"title" url:"title,omitempty"`                                   // title of the outcome.Example: Outcome title
	DisplayName          string          `json:"display_name" url:"display_name,omitempty"`                     // Optional friendly name for reporting.Example: My Favorite Outcome
	Description          string          `json:"description" url:"description,omitempty"`                       // description of the outcome. omitted in the abbreviated form..Example: Outcome description
	VendorGuid           string          `json:"vendor_guid" url:"vendor_guid,omitempty"`                       // A custom GUID for the learning standard..Example: customid9000
	PointsPossible       float64         `json:"points_possible" url:"points_possible,omitempty"`               // maximum points possible. included only if the outcome embeds a rubric criterion. omitted in the abbreviated form..Example: 5
	MasteryPoints        int64           `json:"mastery_points" url:"mastery_points,omitempty"`                 // points necessary to demonstrate mastery outcomes. included only if the outcome embeds a rubric criterion. omitted in the abbreviated form..Example: 3
	CalculationMethod    string          `json:"calculation_method" url:"calculation_method,omitempty"`         // the method used to calculate a students score.Example: decaying_average
	CalculationInt       int64           `json:"calculation_int" url:"calculation_int,omitempty"`               // this defines the variable value used by the calculation_method. included only if calculation_method uses it.Example: 65
	Ratings              []*RubricRating `json:"ratings" url:"ratings,omitempty"`                               // possible ratings for this outcome. included only if the outcome embeds a rubric criterion. omitted in the abbreviated form..
	CanEdit              bool            `json:"can_edit" url:"can_edit,omitempty"`                             // whether the current user can update the outcome.Example: true
	CanUnlink            bool            `json:"can_unlink" url:"can_unlink,omitempty"`                         // whether the outcome can be unlinked.Example: true
	Assessed             bool            `json:"assessed" url:"assessed,omitempty"`                             // whether this outcome has been used to assess a student.Example: true
	HasUpdateableRubrics bool            `json:"has_updateable_rubrics" url:"has_updateable_rubrics,omitempty"` // whether updates to this outcome will propagate to unassessed rubrics that have imported it.Example: true
}

func (t *Outcome) HasErrors() error {
	var s []string
	errs := []string{}
	s = []string{"decaying_average", "n_mastery", "latest", "highest"}
	if t.CalculationMethod != "" && !string_utils.Include(s, t.CalculationMethod) {
		errs = append(errs, fmt.Sprintf("expected 'CalculationMethod' to be one of %v", s))
	}
	return nil
}
