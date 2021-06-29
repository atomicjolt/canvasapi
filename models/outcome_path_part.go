package models

type OutcomePathPart struct {
	Name string `json:"name" url:"name,omitempty"` // The title of the outcome or outcome group.Example: Spelling out numbers
}

func (t *OutcomePathPart) HasErrors() error {
	return nil
}
