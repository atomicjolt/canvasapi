package models

type AnonymousUserDisplay struct {
	AnonymousID    string `json:"anonymous_id" url:"anonymous_id,omitempty"`         // A unique short ID identifying this user within the scope of a particular assignment..Example: xn29Q
	AvatarImageUrl string `json:"avatar_image_url" url:"avatar_image_url,omitempty"` // A URL to retrieve a generic avatar..Example: https://en.gravatar.com/avatar/d8cb8c8cd40ddf0cd05241443a591868?s=80&r=g
}

func (t *AnonymousUserDisplay) HasError() error {
	return nil
}
