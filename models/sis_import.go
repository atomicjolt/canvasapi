package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type SISImport struct {
	ID            int64     `json:"id"`             // The unique identifier for the SIS import..Example: 1
	CreatedAt     time.Time `json:"created_at"`     // The date the SIS import was created..Example: 2013-12-01T23:59:00-06:00
	EndedAt       time.Time `json:"ended_at"`       // The date the SIS import finished. Returns null if not finished..Example: 2013-12-02T00:03:21-06:00
	UpdatedAt     time.Time `json:"updated_at"`     // The date the SIS import was last updated..Example: 2013-12-02T00:03:21-06:00
	WorkflowState string    `json:"workflow_state"` // The current state of the SIS import.
	// - 'initializing': The SIS import is being created, if this gets stuck in initializing, it will not import and will continue on to next import.
	// - 'created': The SIS import has been created.
	// - 'importing': The SIS import is currently processing.
	// - 'cleanup_batch': The SIS import is currently cleaning up courses, sections, and enrollments not included in the batch for batch_mode imports.
	// - 'imported': The SIS import has completed successfully.
	// - 'imported_with_messages': The SIS import completed with errors or warnings.
	// - 'aborted': The SIS import was aborted.
	// - 'failed_with_messages': The SIS import failed with errors.
	// - 'failed': The SIS import failed.
	// - 'restoring': The SIS import is restoring states of imported items.
	// - 'partially_restored': The SIS import is restored some of the states of imported items. This is generally due to passing a param like undelete only.
	// - 'restored': The SIS import is restored all of the states of imported items..Example: imported
	Data                     *SISImportData       `json:"data"`                        // data.
	Statistics               *SISImportStatistics `json:"statistics"`                  // statistics.
	Progress                 string               `json:"progress"`                    // The progress of the SIS import. The progress will reset when using batch_mode and have a different progress for the cleanup stage.Example: 100
	ErrorsAttachment         *File                `json:"errors_attachment"`           // The errors_attachment api object of the SIS import. Only available if there are errors or warning and import has completed..
	User                     *User                `json:"user"`                        // The user that initiated the sis_batch. See the Users API for details..
	ProcessingWarnings       string               `json:"processing_warnings"`         // Only imports that are complete will get this data. An array of CSV_file/warning_message pairs..Example: students.csv, user John Doe has already claimed john_doe's requested login information, skipping
	ProcessingErrors         string               `json:"processing_errors"`           // An array of CSV_file/error_message pairs..Example: students.csv, Error while importing CSV. Please contact support.
	BatchMode                bool                 `json:"batch_mode"`                  // Whether the import was run in batch mode..Example: true
	BatchModeTermID          string               `json:"batch_mode_term_id"`          // The term the batch was limited to..Example: 1234
	MultiTermBatchMode       bool                 `json:"multi_term_batch_mode"`       // Enables batch mode against all terms in term file. Requires change_threshold to be set..Example: false
	SkipDeletes              bool                 `json:"skip_deletes"`                // When set the import will skip any deletes..Example: false
	OverrideSISStickiness    bool                 `json:"override_sis_stickiness"`     // Whether UI changes were overridden..Example: false
	AddSISStickiness         bool                 `json:"add_sis_stickiness"`          // Whether stickiness was added to the batch changes..Example: false
	ClearSISStickiness       bool                 `json:"clear_sis_stickiness"`        // Whether stickiness was cleared..Example: false
	DiffingDataSetIDentifier string               `json:"diffing_data_set_identifier"` // The identifier of the data set that this SIS batch diffs against.Example: account-5-enrollments
	DiffedAgainstImportID    int64                `json:"diffed_against_import_id"`    // The ID of the SIS Import that this import was diffed against.Example: 1
	CsvAttachments           string               `json:"csv_attachments"`             // An array of CSV files for processing.
}

func (t *SISImport) HasError() error {
	var s []string
	s = []string{"initializing", "created", "importing", "cleanup_batch", "imported", "imported_with_messages", "aborted", "failed", "failed_with_messages", "restoring", "partially_restored", "restored"}
	if !string_utils.Include(s, t.WorkflowState) {
		return fmt.Errorf("expected 'workflow_state' to be one of %v", s)
	}
	return nil
}
