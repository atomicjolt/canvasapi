package models

import (
	"time"
)

type LtiAssignment struct {
	ID             int64     `json:"id" url:"id,omitempty"`                           // Example: 4
	Name           string    `json:"name" url:"name,omitempty"`                       // Example: Midterm Review
	Description    string    `json:"description" url:"description,omitempty"`         // Example: <p>Do the following:</p>.
	PointsPossible int64     `json:"points_possible" url:"points_possible,omitempty"` // Example: 10
	DueAt          time.Time `json:"due_at" url:"due_at,omitempty"`                   // The due date for the assignment. If a user id is supplied and an assignment override is in place this field will reflect the due date as it applies to the user..Example: 2012-07-01T23:59:00-06:00
	LtiID          string    `json:"lti_id" url:"lti_id,omitempty"`                   // Example: 86157096483e6b3a50bfedc6bac902c0b20a824f
	CourseID       int64     `json:"course_id" url:"course_id,omitempty"`             // Example: 10000000000060
	LtiCourseID    string    `json:"lti_course_id" url:"lti_course_id,omitempty"`     // Example: 66157096483e6b3a50bfedc6bac902c0b20a8241
}

func (t *LtiAssignment) HasErrors() error {
	return nil
}
