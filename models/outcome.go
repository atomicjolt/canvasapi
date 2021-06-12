package models

import (
	"fmt"

	"github.com/atomicjolt/string_utils"
)

type Outcome struct {
	ID                   int64           `json:"id"`                     // the ID of the outcome.Example: 1
	Url                  string          `json:"url"`                    // the URL for fetching/updating the outcome. should be treated as opaque.Example: /api/v1/outcomes/1
	ContextID            int64           `json:"context_id"`             // the context owning the outcome. may be null for global outcomes.Example: 1
	ContextType          string          `json:"context_type"`           // Example: Account
	Title                string          `json:"title"`                  // title of the outcome.Example: Outcome title
	DisplayName          string          `json:"display_name"`           // Optional friendly name for reporting.Example: My Favorite Outcome
	Description          string          `json:"description"`            // description of the outcome. omitted in the abbreviated form..Example: Outcome description
	VendorGuid           string          `json:"vendor_guid"`            // A custom GUID for the learning standard..Example: customid9000
	PointsPossible       int64           `json:"points_possible"`        // maximum points possible. included only if the outcome embeds a rubric criterion. omitted in the abbreviated form..Example: 5
	MasteryPoints        int64           `json:"mastery_points"`         // points necessary to demonstrate mastery outcomes. included only if the outcome embeds a rubric criterion. omitted in the abbreviated form..Example: 3
	CalculationMethod    string          `json:"calculation_method"`     // the method used to calculate a students score.Example: decaying_average
	CalculationInt       int64           `json:"calculation_int"`        // this defines the variable value used by the calculation_method. included only if calculation_method uses it.Example: 65
	Ratings              []*RubricRating `json:"ratings"`                // possible ratings for this outcome. included only if the outcome embeds a rubric criterion. omitted in the abbreviated form..
	CanEdit              bool            `json:"can_edit"`               // whether the current user can update the outcome.Example: true
	CanUnlink            bool            `json:"can_unlink"`             // whether the outcome can be unlinked.Example: true
	Assessed             bool            `json:"assessed"`               // whether this outcome has been used to assess a student.Example: true
	HasUpdateableRubrics bool            `json:"has_updateable_rubrics"` // whether updates to this outcome will propagate to unassessed rubrics that have imported it.Example: true
}

func (t *Outcome) HasError() error {
	var s []string
	s = []string{"decaying_average", "n_mastery", "latest", "highest"}
	if !string_utils.Include(s, t.CalculationMethod) {
		return fmt.Errorf("expected 'calculation_method' to be one of %v", s)
	}
	return nil
}
