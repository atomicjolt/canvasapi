package models

type NamesAndRoleContext struct {
	ID    string `json:"id" url:"id,omitempty"`       // LTI Context unique identifier.Example: 4dde05e8ca1973bcca9bffc13e1548820eee93a3
	Label string `json:"label" url:"label,omitempty"` // LTI Context short name or code.Example: CS-101
	Title string `json:"title" url:"title,omitempty"` // LTI Context full name.Example: Computer Science 101
}

func (t *NamesAndRoleContext) HasError() error {
	return nil
}
