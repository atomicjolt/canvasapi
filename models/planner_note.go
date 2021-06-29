package models

import (
	"time"
)

type PlannerNote struct {
	ID                  int64     `json:"id" url:"id,omitempty"`                                         // The ID of the planner note.Example: 234
	Title               string    `json:"title" url:"title,omitempty"`                                   // The title for a planner note.Example: Bring books tomorrow
	Description         string    `json:"description" url:"description,omitempty"`                       // The description of the planner note.Example: I need to bring books tomorrow for my course on biology
	UserID              int64     `json:"user_id" url:"user_id,omitempty"`                               // The id of the associated user creating the planner note.Example: 1578941
	WorkflowState       string    `json:"workflow_state" url:"workflow_state,omitempty"`                 // The current published state of the planner note.Example: active
	CourseID            int64     `json:"course_id" url:"course_id,omitempty"`                           // The course that the note is in relation too, if applicable.Example: 1578941
	TodoDate            time.Time `json:"todo_date" url:"todo_date,omitempty"`                           // The datetime of when the planner note should show up on their planner.Example: 2017-05-09T10:12:00Z
	LinkedObjectType    string    `json:"linked_object_type" url:"linked_object_type,omitempty"`         // the type of the linked learning object.Example: assignment
	LinkedObjectID      int64     `json:"linked_object_id" url:"linked_object_id,omitempty"`             // the id of the linked learning object.Example: 131072
	LinkedObjectHtmlUrl string    `json:"linked_object_html_url" url:"linked_object_html_url,omitempty"` // the Canvas web URL of the linked learning object.Example: https://canvas.example.com/courses/1578941/assignments/131072
	LinkedObjectUrl     string    `json:"linked_object_url" url:"linked_object_url,omitempty"`           // the API URL of the linked learning object.Example: https://canvas.example.com/api/v1/courses/1578941/assignments/131072
}

func (t *PlannerNote) HasErrors() error {
	return nil
}
