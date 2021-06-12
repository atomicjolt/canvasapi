package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type SISAssignment struct {
	ID                  int64                               `json:"id"`                     // The unique identifier for the assignment..Example: 4
	CourseID            int64                               `json:"course_id"`              // The unique identifier for the course..Example: 6
	Name                string                              `json:"name"`                   // the name of the assignment.Example: some assignment
	CreatedAt           time.Time                           `json:"created_at"`             // The time at which this assignment was originally created.Example: 2012-07-01T23:59:00-06:00
	DueAt               time.Time                           `json:"due_at"`                 // the due date for the assignment. returns null if not present. NOTE: If this assignment has assignment overrides, this field will be the due date as it applies to the user requesting information from the API..Example: 2012-07-01T23:59:00-06:00
	UnlockAt            time.Time                           `json:"unlock_at"`              // (Optional) Time at which this was/will be unlocked..Example: 2013-01-01T00:00:00-06:00
	LockAt              time.Time                           `json:"lock_at"`                // (Optional) Time at which this was/will be locked..Example: 2013-02-01T00:00:00-06:00
	PointsPossible      int64                               `json:"points_possible"`        // The maximum points possible for the assignment.Example: 12
	SubmissionTypes     string                              `json:"submission_types"`       // the types of submissions allowed for this assignment list containing one or more of the following: 'discussion_topic', 'online_quiz', 'on_paper', 'none', 'external_tool', 'online_text_entry', 'online_url', 'online_upload', 'media_recording', 'student_annotation'.Example: online_text_entry
	IntegrationID       string                              `json:"integration_id"`         // Third Party integration id for assignment.Example: 12341234
	IntegrationData     string                              `json:"integration_data"`       // (optional, Third Party integration data for assignment).Example: other_data
	IncludeInFinalGrade bool                                `json:"include_in_final_grade"` // If false, the assignment will be omitted from the student's final grade.Example: true
	AssignmentGroup     []*AssignmentGroupAttributes        `json:"assignment_group"`       // Includes attributes of a assignment_group for convenience. For more details see Assignments API..
	Sections            []*SectionAttributes                `json:"sections"`               // Includes attributes of a section for convenience. For more details see Sections API..
	UserOverrides       []*UserAssignmentOverrideAttributes `json:"user_overrides"`         // Includes attributes of a user assignment overrides. For more details see Assignments API..
}

func (t *SISAssignment) HasError() error {
	var s []string
	s = []string{"discussion_topic", "online_quiz", "on_paper", "not_graded", "none", "external_tool", "online_text_entry", "online_url", "online_upload", "media_recording", "student_annotation"}
	if !string_utils.Include(s, t.SubmissionTypes) {
		return fmt.Errorf("expected 'submission_types' to be one of %v", s)
	}
	return nil
}
