package models

type QuizIPFilter struct {
	Name    string `json:"name"`    // A unique name for the filter..Example: Current Filter
	Account string `json:"account"` // Name of the Account (or Quiz) the IP filter is defined in..Example: Some Quiz
	Filter  string `json:"filter"`  // An IP address (or range mask) this filter embodies..Example: 192.168.1.1/24
}

func (t *QuizIPFilter) HasError() error {
	return nil
}
