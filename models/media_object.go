package models

type MediaObject struct {
	CanAddCaptions   bool   `json:"can_add_captions" url:"can_add_captions,omitempty"`     //
	UserEnteredTitle string `json:"user_entered_title" url:"user_entered_title,omitempty"` //
	Title            string `json:"title" url:"title,omitempty"`                           //
	MediaID          string `json:"media_id" url:"media_id,omitempty"`                     //
	MediaType        string `json:"media_type" url:"media_type,omitempty"`                 //
	MediaTracks      string `json:"media_tracks" url:"media_tracks,omitempty"`             //
	MediaSources     string `json:"media_sources" url:"media_sources,omitempty"`           //
}

func (t *MediaObject) HasErrors() error {
	return nil
}
