package models

type SectionAttributes struct {
	ID            int64                                `json:"id"`             // The unique identifier for the section..Example: 1
	Name          string                               `json:"name"`           // The name of the section..Example: Section A
	SISID         string                               `json:"sis_id"`         // The sis id of the section..Example: s34643
	IntegrationID string                               `json:"integration_id"` // Optional: The integration ID of the section..Example: 3452342345
	OriginCourse  *CourseAttributes                    `json:"origin_course"`  // The course to which the section belongs or the course from which the section was cross-listed.
	XlistCourse   *CourseAttributes                    `json:"xlist_course"`   // Optional: Attributes of the xlist course. Only present when the section has been cross-listed. See Courses API for more details.
	Override      *SectionAssignmentOverrideAttributes `json:"override"`       // Optional: Attributes of the assignment override that apply to the section. See Assignment API for more details.
}

func (t *SectionAttributes) HasError() error {
	return nil
}
