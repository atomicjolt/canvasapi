package models

type CourseAttributes struct {
	ID            int64  `json:"id"`             // The unique Canvas identifier for the origin course.Example: 7
	Name          string `json:"name"`           // The name of the origin course..Example: Section A
	SISID         string `json:"sis_id"`         // The sis id of the origin_course..Example: c34643
	IntegrationID string `json:"integration_id"` // The integration ID of the origin_course..Example: I-2
}

func (t *CourseAttributes) HasError() error {
	return nil
}
