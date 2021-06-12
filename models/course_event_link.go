package models

type CourseEventLink struct {
	Course     int64  `json:"course"`      // ID of the course for the event..Example: 12345
	User       int64  `json:"user"`        // ID of the user for the event (who made the change)..Example: 12345
	PageView   string `json:"page_view"`   // ID of the page view during the event if it exists..Example: e2b76430-27a5-0131-3ca1-48e0eb13f29b
	CopiedFrom int64  `json:"copied_from"` // ID of the course that this course was copied from. This is only included if the event_type is copied_from..Example: 12345
	CopiedTo   int64  `json:"copied_to"`   // ID of the course that this course was copied to. This is only included if the event_type is copied_to..Example: 12345
	SISBatch   int64  `json:"sis_batch"`   // ID of the SIS batch that triggered the event..Example: 12345
}

func (t *CourseEventLink) HasError() error {
	return nil
}
