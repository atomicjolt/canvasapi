package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type ContentMigration struct {
	ID                 int64     `json:"id"`                   // the unique identifier for the migration.Example: 370663
	MigrationType      string    `json:"migration_type"`       // the type of content migration.Example: common_cartridge_importer
	MigrationTypeTitle string    `json:"migration_type_title"` // the name of the content migration type.Example: Canvas Cartridge Importer
	MigrationIssuesUrl string    `json:"migration_issues_url"` // API url to the content migration's issues.Example: https://example.com/api/v1/courses/1/content_migrations/1/migration_issues
	Attachment         string    `json:"attachment"`           // attachment api object for the uploaded file may not be present for all migrations.Example: {'url'=>'https://example.com/api/v1/courses/1/content_migrations/1/download_archive'}
	ProgressUrl        string    `json:"progress_url"`         // The api endpoint for polling the current progress.Example: https://example.com/api/v1/progress/4
	UserID             int64     `json:"user_id"`              // The user who started the migration.Example: 4
	WorkflowState      string    `json:"workflow_state"`       // Current state of the content migration: pre_processing, pre_processed, running, waiting_for_select, completed, failed.Example: running
	StartedAt          time.Time `json:"started_at"`           // timestamp.Example: 2012-06-01T00:00:00-06:00
	FinishedAt         time.Time `json:"finished_at"`          // timestamp.Example: 2012-06-01T00:00:00-06:00
	PreAttachment      string    `json:"pre_attachment"`       // file uploading data, see {file:file_uploads.html File Upload Documentation} for file upload workflow This works a little differently in that all the file data is in the pre_attachment hash if there is no upload_url then there was an attachment pre-processing error, the error message will be in the message key This data will only be here after a create or update call.Example: {'upload_url'=>'', 'message'=>'file exceeded quota', 'upload_params'=>{}}
}

func (t *ContentMigration) HasError() error {
	var s []string
	s = []string{"pre_processing", "pre_processed", "running", "waiting_for_select", "completed", "failed"}
	if !string_utils.Include(s, t.WorkflowState) {
		return fmt.Errorf("expected 'workflow_state' to be one of %v", s)
	}
	return nil
}
