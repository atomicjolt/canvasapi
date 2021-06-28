package models

type ConversationParticipant struct {
	ID        int64  `json:"id" url:"id,omitempty"`                 // The user ID for the participant..Example: 2
	Name      string `json:"name" url:"name,omitempty"`             // A short name the user has selected, for use in conversations or other less formal places through the site..Example: Shelly
	FullName  string `json:"full_name" url:"full_name,omitempty"`   // The full name of the user..Example: Sheldon Cooper
	AvatarUrl string `json:"avatar_url" url:"avatar_url,omitempty"` // If requested, this field will be included and contain a url to retrieve the user's avatar..Example: https://canvas.instructure.com/images/messages/avatar-50.png
}

func (t *ConversationParticipant) HasError() error {
	return nil
}
