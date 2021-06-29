package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type SISAssignment struct {
	ID                  int64                               `json:"id" url:"id,omitempty"`                                         // The unique identifier for the assignment..Example: 4
	CourseID            int64                               `json:"course_id" url:"course_id,omitempty"`                           // The unique identifier for the course..Example: 6
	Name                string                              `json:"name" url:"name,omitempty"`                                     // the name of the assignment.Example: some assignment
	CreatedAt           time.Time                           `json:"created_at" url:"created_at,omitempty"`                         // The time at which this assignment was originally created.Example: 2012-07-01T23:59:00-06:00
	DueAt               time.Time                           `json:"due_at" url:"due_at,omitempty"`                                 // the due date for the assignment. returns null if not present. NOTE: If this assignment has assignment overrides, this field will be the due date as it applies to the user requesting information from the API..Example: 2012-07-01T23:59:00-06:00
	UnlockAt            time.Time                           `json:"unlock_at" url:"unlock_at,omitempty"`                           // (Optional) Time at which this was/will be unlocked..Example: 2013-01-01T00:00:00-06:00
	LockAt              time.Time                           `json:"lock_at" url:"lock_at,omitempty"`                               // (Optional) Time at which this was/will be locked..Example: 2013-02-01T00:00:00-06:00
	PointsPossible      int64                               `json:"points_possible" url:"points_possible,omitempty"`               // The maximum points possible for the assignment.Example: 12
	SubmissionTypes     []string                            `json:"submission_types" url:"submission_types,omitempty"`             // the types of submissions allowed for this assignment list containing one or more of the following: 'discussion_topic', 'online_quiz', 'on_paper', 'none', 'external_tool', 'online_text_entry', 'online_url', 'online_upload', 'media_recording', 'student_annotation'.Example: online_text_entry
	IntegrationID       string                              `json:"integration_id" url:"integration_id,omitempty"`                 // Third Party integration id for assignment.Example: 12341234
	IntegrationData     string                              `json:"integration_data" url:"integration_data,omitempty"`             // (optional, Third Party integration data for assignment).Example: other_data
	IncludeInFinalGrade bool                                `json:"include_in_final_grade" url:"include_in_final_grade,omitempty"` // If false, the assignment will be omitted from the student's final grade.Example: true
	AssignmentGroup     []*AssignmentGroupAttributes        `json:"assignment_group" url:"assignment_group,omitempty"`             // Includes attributes of a assignment_group for convenience. For more details see Assignments API..
	Sections            []*SectionAttributes                `json:"sections" url:"sections,omitempty"`                             // Includes attributes of a section for convenience. For more details see Sections API..
	UserOverrides       []*UserAssignmentOverrideAttributes `json:"user_overrides" url:"user_overrides,omitempty"`                 // Includes attributes of a user assignment overrides. For more details see Assignments API..
}

func (t *SISAssignment) HasErrors() error {
	var s []string
	errs := []string{}
	s = []string{"discussion_topic", "online_quiz", "on_paper", "not_graded", "none", "external_tool", "online_text_entry", "online_url", "online_upload", "media_recording", "student_annotation"}

	for _, v := range t.SubmissionTypes {
		if v != "" && !string_utils.Include(s, v) {
			errs = append(errs, fmt.Sprintf("expected 'SubmissionTypes' to be one of %v", s))
		}
	}
	return nil
}
