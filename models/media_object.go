package models

type MediaObject struct {
	CanAddCaptions   bool   `json:"can_add_captions"`   //
	UserEnteredTitle string `json:"user_entered_title"` //
	Title            string `json:"title"`              //
	MediaID          string `json:"media_id"`           //
	MediaType        string `json:"media_type"`         //
	MediaTracks      string `json:"media_tracks"`       //
	MediaSources     string `json:"media_sources"`      //
}

func (t *MediaObject) HasError() error {
	return nil
}
