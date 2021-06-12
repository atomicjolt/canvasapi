package models

type Profile struct {
	ID           int64         `json:"id"`            // The ID of the user..Example: 1234
	Name         string        `json:"name"`          // Sample User.Example: Sample User
	ShortName    string        `json:"short_name"`    // Sample User.Example: Sample User
	SortableName string        `json:"sortable_name"` // user, sample.Example: user, sample
	Title        string        `json:"title"`         //
	Bio          string        `json:"bio"`           //
	PrimaryEmail string        `json:"primary_email"` // sample_user@example.com.Example: sample_user@example.com
	LoginID      string        `json:"login_id"`      // sample_user@example.com.Example: sample_user@example.com
	SISUserID    string        `json:"sis_user_id"`   // sis1.Example: sis1
	LtiUserID    string        `json:"lti_user_id"`   //
	AvatarUrl    string        `json:"avatar_url"`    // The avatar_url can change over time, so we recommend not caching it for more than a few hours.Example: url
	Calendar     *CalendarLink `json:"calendar"`      //
	TimeZone     string        `json:"time_zone"`     // Optional: This field is only returned in certain API calls, and will return the IANA time zone name of the user's preferred timezone..Example: America/Denver
	Locale       string        `json:"locale"`        // The users locale..
	K5User       bool          `json:"k5_user"`       // Optional: Whether or not the user is a K5 user. This field is nil if the user settings are not for the user making the request..Example: true
}

func (t *Profile) HasError() error {
	return nil
}
