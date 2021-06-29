package models

type CourseEpubExport struct {
	ID         int64       `json:"id" url:"id,omitempty"`                   // the unique identifier for the course.Example: 101
	Name       string      `json:"name" url:"name,omitempty"`               // the name for the course.Example: Maths 101
	EpubExport *EpubExport `json:"epub_export" url:"epub_export,omitempty"` // ePub export API object.
}

func (t *CourseEpubExport) HasErrors() error {
	return nil
}
