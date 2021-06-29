package models

import (
	"time"
)

type OriginalityReport struct {
	ID                      int64        `json:"id" url:"id,omitempty"`                                                 // The id of the OriginalityReport.Example: 4
	FileID                  int64        `json:"file_id" url:"file_id,omitempty"`                                       // The id of the file receiving the originality score.Example: 8
	OriginalityScore        float64      `json:"originality_score" url:"originality_score,omitempty"`                   // A number between 0 and 100 representing the originality score.Example: 0.16
	OriginalityReportFileID int64        `json:"originality_report_file_id" url:"originality_report_file_id,omitempty"` // The ID of the file within Canvas containing the originality report document (if provided).Example: 23
	OriginalityReportUrl    string       `json:"originality_report_url" url:"originality_report_url,omitempty"`         // A non-LTI launch URL where the originality score of the file may be found..Example: http://www.example.com/report
	ToolSetting             *ToolSetting `json:"tool_setting" url:"tool_setting,omitempty"`                             // A ToolSetting object containing optional 'resource_type_code' and 'resource_url'.
	ErrorReport             string       `json:"error_report" url:"error_report,omitempty"`                             // A message describing the error. If set, the workflow_state will become 'error.'.
	SubmissionTime          time.Time    `json:"submission_time" url:"submission_time,omitempty"`                       // The submitted_at date time of the submission..
	RootAccountID           int64        `json:"root_account_id" url:"root_account_id,omitempty"`                       // The id of the root Account associated with the OriginalityReport.Example: 1
}

func (t *OriginalityReport) HasErrors() error {
	return nil
}
