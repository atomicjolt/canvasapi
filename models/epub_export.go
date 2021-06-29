package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type EpubExport struct {
	ID            int64     `json:"id" url:"id,omitempty"`                         // the unique identifier for the export.Example: 101
	CreatedAt     time.Time `json:"created_at" url:"created_at,omitempty"`         // the date and time this export was requested.Example: 2014-01-01T00:00:00Z
	Attachment    *File     `json:"attachment" url:"attachment,omitempty"`         // attachment api object for the export ePub (not present until the export completes).Example: https://example.com/api/v1/attachments/789?download_frd=1&verifier=bG9sY2F0cyEh
	ProgressUrl   string    `json:"progress_url" url:"progress_url,omitempty"`     // The api endpoint for polling the current progress.Example: https://example.com/api/v1/progress/4
	UserID        int64     `json:"user_id" url:"user_id,omitempty"`               // The ID of the user who started the export.Example: 4
	WorkflowState string    `json:"workflow_state" url:"workflow_state,omitempty"` // Current state of the ePub export: created exporting exported generating generated failed.Example: exported
}

func (t *EpubExport) HasErrors() error {
	var s []string
	errs := []string{}
	s = []string{"created", "exporting", "exported", "generating", "generated", "failed"}
	if t.WorkflowState != "" && !string_utils.Include(s, t.WorkflowState) {
		errs = append(errs, fmt.Sprintf("expected 'WorkflowState' to be one of %v", s))
	}
	return nil
}
