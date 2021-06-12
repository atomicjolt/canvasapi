package models

type MediaTrack struct {
	ID            int64  `json:"id"`              //
	UserID        int64  `json:"user_id"`         //
	MediaObjectID int64  `json:"media_object_id"` //
	Kind          string `json:"kind"`            //
	Locale        string `json:"locale"`          //
	Content       string `json:"content"`         //
	CreatedAt     string `json:"created_at"`      //
	UpdatedAt     string `json:"updated_at"`      //
	WebvttContent string `json:"webvtt_content"`  //
}

func (t *MediaTrack) HasError() error {
	return nil
}
