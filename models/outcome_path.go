package models

type OutcomePath struct {
	ID    int64            `json:"id" url:"id,omitempty"`       // A unique identifier for this outcome.Example: 42
	Parts *OutcomePathPart `json:"parts" url:"parts,omitempty"` // an array of OutcomePathPart objects.
}

func (t *OutcomePath) HasErrors() error {
	return nil
}
