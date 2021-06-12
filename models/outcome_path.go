package models

type OutcomePath struct {
	ID    int64            `json:"id"`    // A unique identifier for this outcome.Example: 42
	Parts *OutcomePathPart `json:"parts"` // an array of OutcomePathPart objects.
}

func (t *OutcomePath) HasError() error {
	return nil
}
