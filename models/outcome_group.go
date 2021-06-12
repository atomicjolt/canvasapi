package models

type OutcomeGroup struct {
	ID                 int64         `json:"id"`                   // the ID of the outcome group.Example: 1
	Url                string        `json:"url"`                  // the URL for fetching/updating the outcome group. should be treated as opaque.Example: /api/v1/accounts/1/outcome_groups/1
	ParentOutcomeGroup *OutcomeGroup `json:"parent_outcome_group"` // an abbreviated OutcomeGroup object representing the parent group of this outcome group, if any. omitted in the abbreviated form..
	ContextID          int64         `json:"context_id"`           // the context owning the outcome group. may be null for global outcome groups. omitted in the abbreviated form..Example: 1
	ContextType        string        `json:"context_type"`         // Example: Account
	Title              string        `json:"title"`                // title of the outcome group.Example: Outcome group title
	Description        string        `json:"description"`          // description of the outcome group. omitted in the abbreviated form..Example: Outcome group description
	VendorGuid         string        `json:"vendor_guid"`          // A custom GUID for the learning standard..Example: customid9000
	SubgroupsUrl       string        `json:"subgroups_url"`        // the URL for listing/creating subgroups under the outcome group. should be treated as opaque.Example: /api/v1/accounts/1/outcome_groups/1/subgroups
	OutcomesUrl        string        `json:"outcomes_url"`         // the URL for listing/creating outcome links under the outcome group. should be treated as opaque.Example: /api/v1/accounts/1/outcome_groups/1/outcomes
	ImportUrl          string        `json:"import_url"`           // the URL for importing another group into this outcome group. should be treated as opaque. omitted in the abbreviated form..Example: /api/v1/accounts/1/outcome_groups/1/import
	CanEdit            bool          `json:"can_edit"`             // whether the current user can update the outcome group.Example: true
}

func (t *OutcomeGroup) HasError() error {
	return nil
}
