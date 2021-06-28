package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type OutcomeImport struct {
	ID            int64     `json:"id" url:"id,omitempty"`                         // The unique identifier for the outcome import..Example: 1
	CreatedAt     time.Time `json:"created_at" url:"created_at,omitempty"`         // The date the outcome import was created..Example: 2013-12-01T23:59:00-06:00
	EndedAt       time.Time `json:"ended_at" url:"ended_at,omitempty"`             // The date the outcome import finished. Returns null if not finished..Example: 2013-12-02T00:03:21-06:00
	UpdatedAt     time.Time `json:"updated_at" url:"updated_at,omitempty"`         // The date the outcome import was last updated..Example: 2013-12-02T00:03:21-06:00
	WorkflowState string    `json:"workflow_state" url:"workflow_state,omitempty"` // The current state of the outcome import.
	// - 'created': The outcome import has been created.
	// - 'importing': The outcome import is currently processing.
	// - 'succeeded': The outcome import has completed successfully.
	// - 'failed': The outcome import failed..Example: imported
	Data             *OutcomeImportData `json:"data" url:"data,omitempty"`                           // See the OutcomeImportData specification above..
	Progress         string             `json:"progress" url:"progress,omitempty"`                   // The progress of the outcome import..Example: 100
	User             *User              `json:"user" url:"user,omitempty"`                           // The user that initiated the outcome_import. See the Users API for details..
	ProcessingErrors string             `json:"processing_errors" url:"processing_errors,omitempty"` // An array of row number / error message pairs. Returns the first 25 errors..Example: 1, Missing required fields: title
}

func (t *OutcomeImport) HasError() error {
	var s []string
	s = []string{"created", "importing", "succeeded", "failed"}
	if t.WorkflowState != "" && !string_utils.Include(s, t.WorkflowState) {
		return fmt.Errorf("expected 'workflow_state' to be one of %v", s)
	}
	return nil
}
