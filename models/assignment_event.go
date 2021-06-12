package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type AssignmentEvent struct {
	ID                  string              `json:"id"`                   // A synthetic ID for the assignment.Example: assignment_987
	Title               string              `json:"title"`                // The title of the assignment.Example: Essay
	StartAt             time.Time           `json:"start_at"`             // The due_at timestamp of the assignment.Example: 2012-07-19T23:59:00-06:00
	EndAt               time.Time           `json:"end_at"`               // The due_at timestamp of the assignment.Example: 2012-07-19T23:59:00-06:00
	Description         string              `json:"description"`          // The HTML description of the assignment.Example: <b>Write an essay. Whatever you want.</b>
	ContextCode         string              `json:"context_code"`         // the context code of the (course) calendar this assignment belongs to.Example: course_123
	WorkflowState       string              `json:"workflow_state"`       // Current state of the assignment ('published' or 'deleted').Example: published
	Url                 string              `json:"url"`                  // URL for this assignment (note that updating/deleting should be done via the Assignments API).Example: https://example.com/api/v1/calendar_events/assignment_987
	HtmlUrl             string              `json:"html_url"`             // URL for a user to view this assignment.Example: http://example.com/courses/123/assignments/987
	AllDayDate          time.Time           `json:"all_day_date"`         // The due date of this assignment.Example: 2012-07-19
	AllDay              bool                `json:"all_day"`              // Boolean indicating whether this is an all-day event (e.g. assignment due at midnight).Example: true
	CreatedAt           time.Time           `json:"created_at"`           // When the assignment was created.Example: 2012-07-12T10:55:20-06:00
	UpdatedAt           time.Time           `json:"updated_at"`           // When the assignment was last updated.Example: 2012-07-12T10:55:20-06:00
	Assignment          *Assignment         `json:"assignment"`           // The full assignment JSON data (See the Assignments API).
	AssignmentOverrides *AssignmentOverride `json:"assignment_overrides"` // The list of AssignmentOverrides that apply to this event (See the Assignments API). This information is useful for determining which students or sections this assignment-due event applies to..
}

func (t *AssignmentEvent) HasError() error {
	var s []string
	s = []string{"published", "deleted"}
	if !string_utils.Include(s, t.WorkflowState) {
		return fmt.Errorf("expected 'workflow_state' to be one of %v", s)
	}
	return nil
}
