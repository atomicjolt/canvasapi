package models

type HelpLink struct {
	ID          string   `json:"id" url:"id,omitempty"`                     // The ID of the help link.Example: instructor_question
	Text        string   `json:"text" url:"text,omitempty"`                 // The name of the help link.Example: Ask Your Instructor a Question
	Subtext     string   `json:"subtext" url:"subtext,omitempty"`           // The description of the help link.Example: Questions are submitted to your instructor
	Url         string   `json:"url" url:"url,omitempty"`                   // The URL of the help link.Example: #teacher_feedback
	Type        string   `json:"type" url:"type,omitempty"`                 // The type of the help link.Example: default
	AvailableTo []string `json:"available_to" url:"available_to,omitempty"` // The roles that have access to this help link.Example: user, student, teacher, admin, observer, unenrolled
}

func (t *HelpLink) HasError() error {
	return nil
}
