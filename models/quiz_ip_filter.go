package models

type QuizIPFilter struct {
	Name    string `json:"name" url:"name,omitempty"`       // A unique name for the filter..Example: Current Filter
	Account string `json:"account" url:"account,omitempty"` // Name of the Account (or Quiz) the IP filter is defined in..Example: Some Quiz
	Filter  string `json:"filter" url:"filter,omitempty"`   // An IP address (or range mask) this filter embodies..Example: 192.168.1.1/24
}

func (t *QuizIPFilter) HasError() error {
	return nil
}
