package models

type Grade struct {
	HtmlUrl               string `json:"html_url" url:"html_url,omitempty"`                               // The URL to the Canvas web UI page for the user's grades, if this is a student enrollment..
	CurrentGrade          string `json:"current_grade" url:"current_grade,omitempty"`                     // The user's current grade in the class. Only included if user has permissions to view this grade..
	FinalGrade            string `json:"final_grade" url:"final_grade,omitempty"`                         // The user's final grade for the class. Only included if user has permissions to view this grade..
	CurrentScore          string `json:"current_score" url:"current_score,omitempty"`                     // The user's current score in the class. Only included if user has permissions to view this score..
	FinalScore            string `json:"final_score" url:"final_score,omitempty"`                         // The user's final score for the class. Only included if user has permissions to view this score..
	CurrentPoints         int64  `json:"current_points" url:"current_points,omitempty"`                   // The total points the user has earned in the class. Only included if user has permissions to view this score and 'current_points' is passed in the request's 'include' parameter..Example: 150
	UnpostedCurrentGrade  string `json:"unposted_current_grade" url:"unposted_current_grade,omitempty"`   // The user's current grade in the class including muted/unposted assignments. Only included if user has permissions to view this grade, typically teachers, TAs, and admins..
	UnpostedFinalGrade    string `json:"unposted_final_grade" url:"unposted_final_grade,omitempty"`       // The user's final grade for the class including muted/unposted assignments. Only included if user has permissions to view this grade, typically teachers, TAs, and admins...
	UnpostedCurrentScore  string `json:"unposted_current_score" url:"unposted_current_score,omitempty"`   // The user's current score in the class including muted/unposted assignments. Only included if user has permissions to view this score, typically teachers, TAs, and admins...
	UnpostedFinalScore    string `json:"unposted_final_score" url:"unposted_final_score,omitempty"`       // The user's final score for the class including muted/unposted assignments. Only included if user has permissions to view this score, typically teachers, TAs, and admins...
	UnpostedCurrentPoints int64  `json:"unposted_current_points" url:"unposted_current_points,omitempty"` // The total points the user has earned in the class, including muted/unposted assignments. Only included if user has permissions to view this score (typically teachers, TAs, and admins) and 'current_points' is passed in the request's 'include' parameter..Example: 150
}

func (t *Grade) HasErrors() error {
	return nil
}
