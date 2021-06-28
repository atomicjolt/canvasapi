package models

type GradeChangeEventLinks struct {
	Assignment int64  `json:"assignment" url:"assignment,omitempty"` // ID of the assignment associated with the event.Example: 2319
	Course     int64  `json:"course" url:"course,omitempty"`         // ID of the course associated with the event. will match the context_id in the associated assignment if the context type for the assignment is a course.Example: 2319
	Student    int64  `json:"student" url:"student,omitempty"`       // ID of the student associated with the event. will match the user_id in the associated submission..Example: 2319
	Grader     int64  `json:"grader" url:"grader,omitempty"`         // ID of the grader associated with the event. will match the grader_id in the associated submission..Example: 2319
	PageView   string `json:"page_view" url:"page_view,omitempty"`   // ID of the page view during the event if it exists..Example: e2b76430-27a5-0131-3ca1-48e0eb13f29b
}

func (t *GradeChangeEventLinks) HasError() error {
	return nil
}
