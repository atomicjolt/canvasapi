package models

type MediaComment struct {
	Contenttype string `json:"content_type"` // Example: audio/mp4
	DisplayName string `json:"display_name"` // Example: something
	MediaID     string `json:"media_id"`     // Example: 3232
	MediaType   string `json:"media_type"`   // Example: audio
	Url         string `json:"url"`          // Example: http://example.com/media_url
}

func (t *MediaComment) HasError() error {
	return nil
}
