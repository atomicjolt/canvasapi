package models

type MediaComment struct {
	Contenttype string `json:"content_type" url:"content_type,omitempty"` // Example: audio/mp4
	DisplayName string `json:"display_name" url:"display_name,omitempty"` // Example: something
	MediaID     string `json:"media_id" url:"media_id,omitempty"`         // Example: 3232
	MediaType   string `json:"media_type" url:"media_type,omitempty"`     // Example: audio
	Url         string `json:"url" url:"url,omitempty"`                   // Example: http://example.com/media_url
}

func (t *MediaComment) HasErrors() error {
	return nil
}
