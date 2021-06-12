package models

type ExceptionRecord struct {
	CourseID           int64  `json:"course_id"`           // The ID of the associated course.Example: 101
	ConflictingChanges string `json:"conflicting_changes"` // A list of change classes in the associated course's copy of the item that prevented a blueprint change from being applied. One or more of ['content', 'points', 'due_dates', 'availability_dates']..Example: points
}

func (t *ExceptionRecord) HasError() error {
	return nil
}
