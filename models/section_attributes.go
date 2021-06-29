package models

type SectionAttributes struct {
	ID            int64                                `json:"id" url:"id,omitempty"`                         // The unique identifier for the section..Example: 1
	Name          string                               `json:"name" url:"name,omitempty"`                     // The name of the section..Example: Section A
	SISID         string                               `json:"sis_id" url:"sis_id,omitempty"`                 // The sis id of the section..Example: s34643
	IntegrationID string                               `json:"integration_id" url:"integration_id,omitempty"` // Optional: The integration ID of the section..Example: 3452342345
	OriginCourse  *CourseAttributes                    `json:"origin_course" url:"origin_course,omitempty"`   // The course to which the section belongs or the course from which the section was cross-listed.
	XlistCourse   *CourseAttributes                    `json:"xlist_course" url:"xlist_course,omitempty"`     // Optional: Attributes of the xlist course. Only present when the section has been cross-listed. See Courses API for more details.
	Override      *SectionAssignmentOverrideAttributes `json:"override" url:"override,omitempty"`             // Optional: Attributes of the assignment override that apply to the section. See Assignment API for more details.
}

func (t *SectionAttributes) HasErrors() error {
	return nil
}
