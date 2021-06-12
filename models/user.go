package models

type User struct {
	ID            int64         `json:"id"`             // The ID of the user..Example: 2
	Name          string        `json:"name"`           // The name of the user..Example: Sheldon Cooper
	SortableName  string        `json:"sortable_name"`  // The name of the user that is should be used for sorting groups of users, such as in the gradebook..Example: Cooper, Sheldon
	ShortName     string        `json:"short_name"`     // A short name the user has selected, for use in conversations or other less formal places through the site..Example: Shelly
	SISUserID     string        `json:"sis_user_id"`    // The SIS ID associated with the user.  This field is only included if the user came from a SIS import and has permissions to view SIS information..Example: SHEL93921
	SISImportID   int64         `json:"sis_import_id"`  // The id of the SIS import.  This field is only included if the user came from a SIS import and has permissions to manage SIS information..Example: 18
	IntegrationID string        `json:"integration_id"` // The integration_id associated with the user.  This field is only included if the user came from a SIS import and has permissions to view SIS information..Example: ABC59802
	LoginID       string        `json:"login_id"`       // The unique login id for the user.  This is what the user uses to log in to Canvas..Example: sheldon@caltech.example.com
	AvatarUrl     string        `json:"avatar_url"`     // If avatars are enabled, this field will be included and contain a url to retrieve the user's avatar..Example: https://en.gravatar.com/avatar/d8cb8c8cd40ddf0cd05241443a591868?s=80&r=g
	Enrollments   []*Enrollment `json:"enrollments"`    // Optional: This field can be requested with certain API calls, and will return a list of the users active enrollments. See the List enrollments API for more details about the format of these records..
	Email         string        `json:"email"`          // Optional: This field can be requested with certain API calls, and will return the users primary email address..Example: sheldon@caltech.example.com
	Locale        string        `json:"locale"`         // Optional: This field can be requested with certain API calls, and will return the users locale in RFC 5646 format..Example: tlh
	LastLogin     string        `json:"last_login"`     // Optional: This field is only returned in certain API calls, and will return a timestamp representing the last time the user logged in to canvas..Example: 2012-05-30T17:45:25Z
	TimeZone      string        `json:"time_zone"`      // Optional: This field is only returned in certain API calls, and will return the IANA time zone name of the user's preferred timezone..Example: America/Denver
	Bio           string        `json:"bio"`            // Optional: The user's bio..Example: I like the Muppets.
}

func (t *User) HasError() error {
	return nil
}
