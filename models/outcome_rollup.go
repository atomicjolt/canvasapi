package models

type OutcomeRollup struct {
	Scores *OutcomeRollupScore `json:"scores" url:"scores,omitempty"` // an array of OutcomeRollupScore objects.
	Name   string              `json:"name" url:"name,omitempty"`     // The name of the resource for this rollup. For example, the user name..Example: John Doe
	Links  *OutcomeRollupLinks `json:"links" url:"links,omitempty"`   // Example: 42, 42, 57
}

func (t *OutcomeRollup) HasErrors() error {
	return nil
}
