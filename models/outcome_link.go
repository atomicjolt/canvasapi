package models

type OutcomeLink struct {
	Url          string        `json:"url" url:"url,omitempty"`                     // the URL for fetching/updating the outcome link. should be treated as opaque.Example: /api/v1/accounts/1/outcome_groups/1/outcomes/1
	ContextID    int64         `json:"context_id" url:"context_id,omitempty"`       // the context owning the outcome link. will match the context owning the outcome group containing the outcome link; included for convenience. may be null for links in global outcome groups..Example: 1
	ContextType  string        `json:"context_type" url:"context_type,omitempty"`   // Example: Account
	OutcomeGroup *OutcomeGroup `json:"outcome_group" url:"outcome_group,omitempty"` // an abbreviated OutcomeGroup object representing the group containing the outcome link..
	Outcome      *Outcome      `json:"outcome" url:"outcome,omitempty"`             // an abbreviated Outcome object representing the outcome linked into the containing outcome group..
	Assessed     bool          `json:"assessed" url:"assessed,omitempty"`           // whether this outcome has been used to assess a student in the context of this outcome link.  In other words, this will be set to true if the context is a course, and a student has been assessed with this outcome in that course..Example: true
	CanUnlink    bool          `json:"can_unlink" url:"can_unlink,omitempty"`       // whether this outcome link is manageable and is not the last link to an aligned outcome.
}

func (t *OutcomeLink) HasError() error {
	return nil
}
