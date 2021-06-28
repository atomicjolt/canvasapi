package models

type Avatar struct {
	Type        string `json:"type" url:"type,omitempty"`                 // ['gravatar'|'attachment'|'no_pic'] The type of avatar record, for categorization purposes..Example: gravatar
	Url         string `json:"url" url:"url,omitempty"`                   // The url of the avatar.Example: https://secure.gravatar.com/avatar/2284.
	Token       string `json:"token" url:"token,omitempty"`               // A unique representation of the avatar record which can be used to set the avatar with the user update endpoint. Note: this is an internal representation and is subject to change without notice. It should be consumed with this api endpoint and used in the user update endpoint, and should not be constructed by the client..Example: <opaque_token>
	DisplayName string `json:"display_name" url:"display_name,omitempty"` // A textual description of the avatar record..Example: user, sample
	ID          int64  `json:"id" url:"id,omitempty"`                     // ['attachment' type only] the internal id of the attachment.Example: 12
	Contenttype string `json:"content_type" url:"content_type,omitempty"` // ['attachment' type only] the content-type of the attachment..Example: image/jpeg
	Filename    string `json:"filename" url:"filename,omitempty"`         // ['attachment' type only] the filename of the attachment.Example: profile.jpg
	Size        int64  `json:"size" url:"size,omitempty"`                 // ['attachment' type only] the size of the attachment.Example: 32649
}

func (t *Avatar) HasError() error {
	return nil
}
