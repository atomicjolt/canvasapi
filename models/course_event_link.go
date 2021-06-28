package models

type CourseEventLink struct {
	Course     int64  `json:"course" url:"course,omitempty"`           // ID of the course for the event..Example: 12345
	User       int64  `json:"user" url:"user,omitempty"`               // ID of the user for the event (who made the change)..Example: 12345
	PageView   string `json:"page_view" url:"page_view,omitempty"`     // ID of the page view during the event if it exists..Example: e2b76430-27a5-0131-3ca1-48e0eb13f29b
	CopiedFrom int64  `json:"copied_from" url:"copied_from,omitempty"` // ID of the course that this course was copied from. This is only included if the event_type is copied_from..Example: 12345
	CopiedTo   int64  `json:"copied_to" url:"copied_to,omitempty"`     // ID of the course that this course was copied to. This is only included if the event_type is copied_to..Example: 12345
	SISBatch   int64  `json:"sis_batch" url:"sis_batch,omitempty"`     // ID of the SIS batch that triggered the event..Example: 12345
}

func (t *CourseEventLink) HasError() error {
	return nil
}
