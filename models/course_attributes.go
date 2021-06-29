package models

type CourseAttributes struct {
	ID            int64  `json:"id" url:"id,omitempty"`                         // The unique Canvas identifier for the origin course.Example: 7
	Name          string `json:"name" url:"name,omitempty"`                     // The name of the origin course..Example: Section A
	SISID         string `json:"sis_id" url:"sis_id,omitempty"`                 // The sis id of the origin_course..Example: c34643
	IntegrationID string `json:"integration_id" url:"integration_id,omitempty"` // The integration ID of the origin_course..Example: I-2
}

func (t *CourseAttributes) HasErrors() error {
	return nil
}
