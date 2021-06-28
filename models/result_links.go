package models

type ResultLinks struct {
	User       string `json:"user" url:"user,omitempty"`             // A unique identifier for the user to whom this result applies.Example: 42
	Assessor   string `json:"assessor" url:"assessor,omitempty"`     // A unique identifier for the user who created this result.Example: 23
	Assessment string `json:"assessment" url:"assessment,omitempty"` // A unique identifier for the assessment that this result is for.Example: 5
}

func (t *ResultLinks) HasError() error {
	return nil
}
