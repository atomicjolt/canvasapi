package models

type CustomColumn struct {
	ID           int64  `json:"id" url:"id,omitempty"`                       // The ID of the custom gradebook column.Example: 2
	TeacherNotes bool   `json:"teacher_notes" url:"teacher_notes,omitempty"` // When true, this column's visibility will be toggled in the Gradebook when a user selects to show or hide notes.
	Title        string `json:"title" url:"title,omitempty"`                 // header text.Example: Stuff
	Position     int64  `json:"position" url:"position,omitempty"`           // column order.Example: 1
	Hidden       bool   `json:"hidden" url:"hidden,omitempty"`               // won't be displayed if hidden is true.
	ReadOnly     bool   `json:"read_only" url:"read_only,omitempty"`         // won't be editable in the gradebook UI.Example: true
}

func (t *CustomColumn) HasErrors() error {
	return nil
}
