package models

type CourseEpubExport struct {
	ID         int64       `json:"id"`          // the unique identifier for the course.Example: 101
	Name       string      `json:"name"`        // the name for the course.Example: Maths 101
	EpubExport *EpubExport `json:"epub_export"` // ePub export API object.
}

func (t *CourseEpubExport) HasError() error {
	return nil
}
