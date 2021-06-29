package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type Progress struct {
	ID            int64                    `json:"id" url:"id,omitempty"`                         // the ID of the Progress object.Example: 1
	ContextID     int64                    `json:"context_id" url:"context_id,omitempty"`         // the context owning the job..Example: 1
	ContextType   string                   `json:"context_type" url:"context_type,omitempty"`     // Example: Account
	UserID        int64                    `json:"user_id" url:"user_id,omitempty"`               // the id of the user who started the job.Example: 123
	Tag           string                   `json:"tag" url:"tag,omitempty"`                       // the type of operation.Example: course_batch_update
	Completion    int64                    `json:"completion" url:"completion,omitempty"`         // percent completed.Example: 100
	WorkflowState string                   `json:"workflow_state" url:"workflow_state,omitempty"` // the state of the job one of 'queued', 'running', 'completed', 'failed'.Example: completed
	CreatedAt     time.Time                `json:"created_at" url:"created_at,omitempty"`         // the time the job was created.Example: 2013-01-15T15:00:00Z
	UpdatedAt     time.Time                `json:"updated_at" url:"updated_at,omitempty"`         // the time the job was last updated.Example: 2013-01-15T15:04:00Z
	Message       string                   `json:"message" url:"message,omitempty"`               // optional details about the job.Example: 17 courses processed
	Results       map[string](interface{}) `json:"results" url:"results,omitempty"`               // optional results of the job. omitted when job is still pending.Example: 123
	Url           string                   `json:"url" url:"url,omitempty"`                       // url where a progress update can be retrieved with an LTI access token.Example: https://canvas.example.edu/api/lti/courses/1/progress/1
}

func (t *Progress) HasErrors() error {
	var s []string
	errs := []string{}
	s = []string{"queued", "running", "completed", "failed"}
	if t.WorkflowState != "" && !string_utils.Include(s, t.WorkflowState) {
		errs = append(errs, fmt.Sprintf("expected 'WorkflowState' to be one of %v", s))
	}
	return nil
}
