package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type ReportParameters struct {
	EnrollmentTermID       int64     `json:"enrollment_term_id"`       // The canvas id of the term to get grades from.Example: 2
	IncludeDeleted         bool      `json:"include_deleted"`          // If true, deleted objects will be included. If false, deleted objects will be omitted..
	CourseID               int64     `json:"course_id"`                // The id of the course to report on.Example: 2
	Order                  string    `json:"order"`                    // The sort order for the csv, Options: 'users', 'courses', 'outcomes'..Example: users
	Users                  bool      `json:"users"`                    // If true, user data will be included. If false, user data will be omitted..
	Accounts               bool      `json:"accounts"`                 // If true, account data will be included. If false, account data will be omitted..
	Terms                  bool      `json:"terms"`                    // If true, term data will be included. If false, term data will be omitted..
	Courses                bool      `json:"courses"`                  // If true, course data will be included. If false, course data will be omitted..
	Sections               bool      `json:"sections"`                 // If true, section data will be included. If false, section data will be omitted..
	Enrollments            bool      `json:"enrollments"`              // If true, enrollment data will be included. If false, enrollment data will be omitted..
	Groups                 bool      `json:"groups"`                   // If true, group data will be included. If false, group data will be omitted..
	Xlist                  bool      `json:"xlist"`                    // If true, data for crosslisted courses will be included. If false, data for crosslisted courses will be omitted..
	SISTermsCsv            int64     `json:"sis_terms_csv"`            // Example: 1
	SISAccountsCsv         int64     `json:"sis_accounts_csv"`         // Example: 1
	IncludeEnrollmentState bool      `json:"include_enrollment_state"` // If true, enrollment state will be included. If false, enrollment state will be omitted. Defaults to false..
	EnrollmentState        []string  `json:"enrollment_state"`         // Include enrollment state. Defaults to 'all' Options: ['active'| 'invited'| 'creation_pending'| 'deleted'| 'rejected'| 'completed'| 'inactive'| 'all'].Example: all
	StartAt                time.Time `json:"start_at"`                 // The beginning date for submissions. Max time range is 2 weeks..Example: 2012-07-13T10:55:20-06:00
	EndAt                  time.Time `json:"end_at"`                   // The end date for submissions. Max time range is 2 weeks..Example: 2012-07-13T10:55:20-06:00
}

func (t *ReportParameters) HasError() error {
	var s []string
	s = []string{"users", "courses", "outcomes"}
	if !string_utils.Include(s, t.Order) {
		return fmt.Errorf("expected 'order' to be one of %v", s)
	}
	return nil
}
