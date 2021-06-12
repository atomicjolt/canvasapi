package models

type OutcomeRollup struct {
	Scores *OutcomeRollupScore `json:"scores"` // an array of OutcomeRollupScore objects.
	Name   string              `json:"name"`   // The name of the resource for this rollup. For example, the user name..Example: John Doe
	Links  *OutcomeRollupLinks `json:"links"`  // Example: 42, 42, 57
}

func (t *OutcomeRollup) HasError() error {
	return nil
}
