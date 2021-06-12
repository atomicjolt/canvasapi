package models

import (
	"time"
)

type Section struct {
	ID                                int64     `json:"id"`                                    // The unique identifier for the section..Example: 1
	Name                              string    `json:"name"`                                  // The name of the section..Example: Section A
	SISSectionID                      string    `json:"sis_section_id"`                        // The sis id of the section. This field is only included if the user has permission to view SIS information..Example: s34643
	IntegrationID                     string    `json:"integration_id"`                        // Optional: The integration ID of the section. This field is only included if the user has permission to view SIS information..Example: 3452342345
	SISImportID                       int64     `json:"sis_import_id"`                         // The unique identifier for the SIS import if created through SIS. This field is only included if the user has permission to manage SIS information..Example: 47
	CourseID                          int64     `json:"course_id"`                             // The unique Canvas identifier for the course in which the section belongs.Example: 7
	SISCourseID                       string    `json:"sis_course_id"`                         // The unique SIS identifier for the course in which the section belongs. This field is only included if the user has permission to view SIS information..Example: 7
	StartAt                           time.Time `json:"start_at"`                              // the start date for the section, if applicable.Example: 2012-06-01T00:00:00-06:00
	EndAt                             time.Time `json:"end_at"`                                // the end date for the section, if applicable.
	RestrictEnrollmentsToSectionDates bool      `json:"restrict_enrollments_to_section_dates"` // Restrict user enrollments to the start and end dates of the section.
	NonxlistCourseID                  int64     `json:"nonxlist_course_id"`                    // The unique identifier of the original course of a cross-listed section.
	TotalStudents                     int64     `json:"total_students"`                        // optional: the total number of active and invited students in the section.Example: 13
}

func (t *Section) HasError() error {
	return nil
}
