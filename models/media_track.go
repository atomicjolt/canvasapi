package models

type MediaTrack struct {
	ID            int64  `json:"id" url:"id,omitempty"`                           //
	UserID        int64  `json:"user_id" url:"user_id,omitempty"`                 //
	MediaObjectID int64  `json:"media_object_id" url:"media_object_id,omitempty"` //
	Kind          string `json:"kind" url:"kind,omitempty"`                       //
	Locale        string `json:"locale" url:"locale,omitempty"`                   //
	Content       string `json:"content" url:"content,omitempty"`                 //
	CreatedAt     string `json:"created_at" url:"created_at,omitempty"`           //
	UpdatedAt     string `json:"updated_at" url:"updated_at,omitempty"`           //
	WebvttContent string `json:"webvtt_content" url:"webvtt_content,omitempty"`   //
}

func (t *MediaTrack) HasError() error {
	return nil
}
