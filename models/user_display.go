package models

type UserDisplay struct {
	ID             int64  `json:"id"`               // The ID of the user..Example: 2
	ShortName      string `json:"short_name"`       // A short name the user has selected, for use in conversations or other less formal places through the site..Example: Shelly
	AvatarImageUrl string `json:"avatar_image_url"` // If avatars are enabled, this field will be included and contain a url to retrieve the user's avatar..Example: https://en.gravatar.com/avatar/d8cb8c8cd40ddf0cd05241443a591868?s=80&r=g
	HtmlUrl        string `json:"html_url"`         // URL to access user, either nested to a context or directly..Example: https://school.instructure.com/courses/:course_id/users/:user_id
}

func (t *UserDisplay) HasError() error {
	return nil
}
