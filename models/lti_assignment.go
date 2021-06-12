package models

import (
	"time"
)

type LtiAssignment struct {
	ID             int64     `json:"id"`              // Example: 4
	Name           string    `json:"name"`            // Example: Midterm Review
	Description    string    `json:"description"`     // Example: <p>Do the following:</p>.
	PointsPossible int64     `json:"points_possible"` // Example: 10
	DueAt          time.Time `json:"due_at"`          // The due date for the assignment. If a user id is supplied and an assignment override is in place this field will reflect the due date as it applies to the user..Example: 2012-07-01T23:59:00-06:00
	LtiID          string    `json:"lti_id"`          // Example: 86157096483e6b3a50bfedc6bac902c0b20a824f
	CourseID       int64     `json:"course_id"`       // Example: 10000000000060
	LtiCourseID    string    `json:"lti_course_id"`   // Example: 66157096483e6b3a50bfedc6bac902c0b20a8241
}

func (t *LtiAssignment) HasError() error {
	return nil
}
