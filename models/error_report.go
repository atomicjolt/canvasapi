package models

type ErrorReport struct {
	Subject               string `json:"subject"`                 // The users problem summary, like an email subject line.Example: File upload breaking
	Comments              string `json:"comments"`                // long form documentation of what was witnessed.Example: When I went to upload a .mov file to my files page, I got an error.  Retrying didn't help, other file types seem ok
	UserPerceivedSeverity string `json:"user_perceived_severity"` // categorization of how bad the user thinks the problem is.  Should be one of [just_a_comment, not_urgent, workaround_possible, blocks_what_i_need_to_do, extreme_critical_emergency]..Example: just_a_comment
	Email                 string `json:"email"`                   // the email address of the reporting user.Example: name@example.com
	Url                   string `json:"url"`                     // URL of the page on which the error was reported.Example: https://canvas.instructure.com/courses/1
	ContextAssetString    string `json:"context_asset_string"`    // string describing the asset being interacted with at the time of error.  Formatted '[type]_[id]'.Example: user_1
	UserRoles             string `json:"user_roles"`              // comma seperated list of roles the reporting user holds.  Can be one [student], or many [teacher,admin].Example: user,teacher,admin
}

func (t *ErrorReport) HasError() error {
	return nil
}
