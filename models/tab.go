package models

type Tab struct {
	HtmlUrl    string `json:"html_url"`   // Example: /courses/1/external_tools/4
	ID         string `json:"id"`         // Example: context_external_tool_4
	Label      string `json:"label"`      // Example: WordPress
	Type       string `json:"type"`       // Example: external
	Hidden     bool   `json:"hidden"`     // only included if true.Example: true
	Visibility string `json:"visibility"` // possible values are: public, members, admins, and none.Example: public
	Position   int64  `json:"position"`   // 1 based.Example: 2
}

func (t *Tab) HasError() error {
	return nil
}
