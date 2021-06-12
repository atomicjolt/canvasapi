package models

type CustomColumn struct {
	ID           int64  `json:"id"`            // The ID of the custom gradebook column.Example: 2
	TeacherNotes bool   `json:"teacher_notes"` // When true, this column's visibility will be toggled in the Gradebook when a user selects to show or hide notes.
	Title        string `json:"title"`         // header text.Example: Stuff
	Position     int64  `json:"position"`      // column order.Example: 1
	Hidden       bool   `json:"hidden"`        // won't be displayed if hidden is true.
	ReadOnly     bool   `json:"read_only"`     // won't be editable in the gradebook UI.Example: true
}

func (t *CustomColumn) HasError() error {
	return nil
}
