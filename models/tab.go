package models

type Tab struct {
	HtmlUrl    string `json:"html_url" url:"html_url,omitempty"`     // Example: /courses/1/external_tools/4
	ID         string `json:"id" url:"id,omitempty"`                 // Example: context_external_tool_4
	Label      string `json:"label" url:"label,omitempty"`           // Example: WordPress
	Type       string `json:"type" url:"type,omitempty"`             // Example: external
	Hidden     bool   `json:"hidden" url:"hidden,omitempty"`         // only included if true.Example: true
	Visibility string `json:"visibility" url:"visibility,omitempty"` // possible values are: public, members, admins, and none.Example: public
	Position   int64  `json:"position" url:"position,omitempty"`     // 1 based.Example: 2
}

func (t *Tab) HasError() error {
	return nil
}
