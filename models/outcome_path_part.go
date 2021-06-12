package models

type OutcomePathPart struct {
	Name string `json:"name"` // The title of the outcome or outcome group.Example: Spelling out numbers
}

func (t *OutcomePathPart) HasError() error {
	return nil
}
