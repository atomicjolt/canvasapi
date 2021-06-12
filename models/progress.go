package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type Progress struct {
	ID            int64     `json:"id"`             // the ID of the Progress object.Example: 1
	ContextID     int64     `json:"context_id"`     // the context owning the job..Example: 1
	ContextType   string    `json:"context_type"`   // Example: Account
	UserID        int64     `json:"user_id"`        // the id of the user who started the job.Example: 123
	Tag           string    `json:"tag"`            // the type of operation.Example: course_batch_update
	Completion    int64     `json:"completion"`     // percent completed.Example: 100
	WorkflowState string    `json:"workflow_state"` // the state of the job one of 'queued', 'running', 'completed', 'failed'.Example: completed
	CreatedAt     time.Time `json:"created_at"`     // the time the job was created.Example: 2013-01-15T15:00:00Z
	UpdatedAt     time.Time `json:"updated_at"`     // the time the job was last updated.Example: 2013-01-15T15:04:00Z
	Message       string    `json:"message"`        // optional details about the job.Example: 17 courses processed
	Results       string    `json:"results"`        // optional results of the job. omitted when job is still pending.Example: 123
	Url           string    `json:"url"`            // url where a progress update can be retrieved with an LTI access token.Example: https://canvas.example.edu/api/lti/courses/1/progress/1
}

func (t *Progress) HasError() error {
	var s []string
	s = []string{"queued", "running", "completed", "failed"}
	if !string_utils.Include(s, t.WorkflowState) {
		return fmt.Errorf("expected 'workflow_state' to be one of %v", s)
	}
	return nil
}
