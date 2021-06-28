package models

type Profile struct {
	ID           int64         `json:"id" url:"id,omitempty"`                       // The ID of the user..Example: 1234
	Name         string        `json:"name" url:"name,omitempty"`                   // Sample User.Example: Sample User
	ShortName    string        `json:"short_name" url:"short_name,omitempty"`       // Sample User.Example: Sample User
	SortableName string        `json:"sortable_name" url:"sortable_name,omitempty"` // user, sample.Example: user, sample
	Title        string        `json:"title" url:"title,omitempty"`                 //
	Bio          string        `json:"bio" url:"bio,omitempty"`                     //
	PrimaryEmail string        `json:"primary_email" url:"primary_email,omitempty"` // sample_user@example.com.Example: sample_user@example.com
	LoginID      string        `json:"login_id" url:"login_id,omitempty"`           // sample_user@example.com.Example: sample_user@example.com
	SISUserID    string        `json:"sis_user_id" url:"sis_user_id,omitempty"`     // sis1.Example: sis1
	LtiUserID    string        `json:"lti_user_id" url:"lti_user_id,omitempty"`     //
	AvatarUrl    string        `json:"avatar_url" url:"avatar_url,omitempty"`       // The avatar_url can change over time, so we recommend not caching it for more than a few hours.Example: url
	Calendar     *CalendarLink `json:"calendar" url:"calendar,omitempty"`           //
	TimeZone     string        `json:"time_zone" url:"time_zone,omitempty"`         // Optional: This field is only returned in certain API calls, and will return the IANA time zone name of the user's preferred timezone..Example: America/Denver
	Locale       string        `json:"locale" url:"locale,omitempty"`               // The users locale..
	K5User       bool          `json:"k5_user" url:"k5_user,omitempty"`             // Optional: Whether or not the user is a K5 user. This field is nil if the user settings are not for the user making the request..Example: true
}

func (t *Profile) HasError() error {
	return nil
}
