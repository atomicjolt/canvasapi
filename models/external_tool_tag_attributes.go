package models

type ExternalToolTagAttributes struct {
	Url            string `json:"url"`              // URL to the external tool.Example: http://instructure.com
	NewTab         bool   `json:"new_tab"`          // Whether or not there is a new tab for the external tool.
	ResourceLinkID string `json:"resource_link_id"` // the identifier for this tool_tag.Example: ab81173af98b8c33e66a
}

func (t *ExternalToolTagAttributes) HasError() error {
	return nil
}
