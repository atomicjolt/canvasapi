package models

type Grade struct {
	HtmlUrl               string `json:"html_url"`                // The URL to the Canvas web UI page for the user's grades, if this is a student enrollment..
	CurrentGrade          string `json:"current_grade"`           // The user's current grade in the class. Only included if user has permissions to view this grade..
	FinalGrade            string `json:"final_grade"`             // The user's final grade for the class. Only included if user has permissions to view this grade..
	CurrentScore          string `json:"current_score"`           // The user's current score in the class. Only included if user has permissions to view this score..
	FinalScore            string `json:"final_score"`             // The user's final score for the class. Only included if user has permissions to view this score..
	CurrentPoints         int64  `json:"current_points"`          // The total points the user has earned in the class. Only included if user has permissions to view this score and 'current_points' is passed in the request's 'include' parameter..Example: 150
	UnpostedCurrentGrade  string `json:"unposted_current_grade"`  // The user's current grade in the class including muted/unposted assignments. Only included if user has permissions to view this grade, typically teachers, TAs, and admins..
	UnpostedFinalGrade    string `json:"unposted_final_grade"`    // The user's final grade for the class including muted/unposted assignments. Only included if user has permissions to view this grade, typically teachers, TAs, and admins...
	UnpostedCurrentScore  string `json:"unposted_current_score"`  // The user's current score in the class including muted/unposted assignments. Only included if user has permissions to view this score, typically teachers, TAs, and admins...
	UnpostedFinalScore    string `json:"unposted_final_score"`    // The user's final score for the class including muted/unposted assignments. Only included if user has permissions to view this score, typically teachers, TAs, and admins...
	UnpostedCurrentPoints int64  `json:"unposted_current_points"` // The total points the user has earned in the class, including muted/unposted assignments. Only included if user has permissions to view this score (typically teachers, TAs, and admins) and 'current_points' is passed in the request's 'include' parameter..Example: 150
}

func (t *Grade) HasError() error {
	return nil
}
