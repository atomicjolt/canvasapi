package models

type ExternalToolTagAttributes struct {
	Url            string `json:"url" url:"url,omitempty"`                           // URL to the external tool.Example: http://instructure.com
	NewTab         bool   `json:"new_tab" url:"new_tab,omitempty"`                   // Whether or not there is a new tab for the external tool.
	ResourceLinkID string `json:"resource_link_id" url:"resource_link_id,omitempty"` // the identifier for this tool_tag.Example: ab81173af98b8c33e66a
}

func (t *ExternalToolTagAttributes) HasErrors() error {
	return nil
}
