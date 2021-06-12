package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type ContentExport struct {
	ID            int64     `json:"id"`             // the unique identifier for the export.Example: 101
	CreatedAt     time.Time `json:"created_at"`     // the date and time this export was requested.Example: 2014-01-01T00:00:00Z
	ExportType    string    `json:"export_type"`    // the type of content migration: 'common_cartridge' or 'qti'.Example: common_cartridge
	Attachment    *File     `json:"attachment"`     // attachment api object for the export package (not present before the export completes or after it becomes unavailable for download.).Example: https://example.com/api/v1/attachments/789?download_frd=1&verifier=bG9sY2F0cyEh
	ProgressUrl   string    `json:"progress_url"`   // The api endpoint for polling the current progress.Example: https://example.com/api/v1/progress/4
	UserID        int64     `json:"user_id"`        // The ID of the user who started the export.Example: 4
	WorkflowState string    `json:"workflow_state"` // Current state of the content migration: created exporting exported failed.Example: exported
}

func (t *ContentExport) HasError() error {
	var s []string
	s = []string{"common_cartridge", "qti"}
	if !string_utils.Include(s, t.ExportType) {
		return fmt.Errorf("expected 'export_type' to be one of %v", s)
	}
	s = []string{"created", "exporting", "exported", "failed"}
	if !string_utils.Include(s, t.WorkflowState) {
		return fmt.Errorf("expected 'workflow_state' to be one of %v", s)
	}
	return nil
}
